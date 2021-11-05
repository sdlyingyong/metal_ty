package models

import (
	"errors"
	"github.com/prometheus/common/log"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/beego/beego/v2/client/orm"
)

//set response json data
type Article struct {
	BaseModel
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   int8   `json:"status"`
	Category string `json:"category"`
	Keywords string `json:"keywords"`
}

//set response json data
type ArticlePortal struct {
	Article
	ViewCount int
	Img       string
	CreatedAt string
	UpdatedAt string
	Previous  Article
	Next      Article
}

func (t *Article) TableName() string {
	return "article"
}

func init() {
	orm.RegisterModel(new(Article))
}

// AddArticle insert a new Article into database and returns
// last inserted Id on success.
func AddArticle(m *Article) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticleById retrieves Article by Id. Returns error if
// Id doesn't exist
//func GetArticleById(id int) (v *Article, err error) {
//	o := orm.NewOrm()
//	v = &Article{Id: id}
//	if err = o.Read(v); err == nil {
//		return v, nil
//	}
//	return nil, err
//}

// GetAllArticle retrieves all Article matches certain condition. Returns empty list if
// no records exist
func GetAllArticle(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Article
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateArticle updates Article by Id and returns error if
// the record to be updated doesn't exist
//func UpdateArticleById(m *Article) (err error) {
//	o := orm.NewOrm()
//	v := Article{Id: m.Id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		if num, err = o.Update(m); err == nil {
//			fmt.Println("Number of records updated in database:", num)
//		}
//	}
//	return
//}

// DeleteArticle deletes Article by Id and returns error if
// the record to be deleted doesn't exist
//func DeleteArticle(id int) (err error) {
//	o := orm.NewOrm()
//	v := Article{Id: id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		if num, err = o.Delete(&Article{Id: id}); err == nil {
//			fmt.Println("Number of records deleted in database:", num)
//		}
//	}
//	return
//}

//get article list and count
//with go routine
func (model *Article) GetArticlesByCondition(param map[string]string, pageIndex, pageSize int) (articles []Article, total int64, returnError error) {
	o := orm.NewOrm()
	var condition = ""
	if param["title"] != "" {
		condition += " AND title LIKE '" + param["title"] + "%'"
	}
	if param["category"] != "" {
		condition += "AND category = '" + param["category"] + "%'"
	}
	if param["keywords"] != "" {
		condition += "AND keywords LIKE '% " + param["keywords"] + "%'"
	}
	var wg sync.WaitGroup
	wg.Add(2)
	//Parallel query 1
	go func() {
		//send execute end message
		defer wg.Done()
		var sql = "SELECT * FROM article WHERE status = 1"
		sql += condition
		sql += " ORDER BY id DESC"
		sql += " LIMIT ?, ?;"
		//save sql result in object
		_, err := o.Raw(sql, pageIndex, pageSize).QueryRows(&articles)
		if err != nil {
			returnError = err
		}
	}()
	//Parallel query 2
	go func() {
		//end message
		defer wg.Done()
		var sql = "SELECT COUNT(0) FROM article WHERE status = 1"
		sql += condition
		err := o.Raw(sql).QueryRow(&total)
		if err != nil {
			returnError = err
		}
		log.Info("mysql row affected nums: ", total)
	}()
	wg.Wait()
	return articles, total, returnError
}

//update one article
func (model *Article) Update() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(model, "title", "content", "category", "keywords", "updated_at")
	return id, err
}

//delete one article
func (model *Article) Delete() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Delete(model)
	return id, err
}

//get one article by id
func (model *Article) GetById() error {
	o := orm.NewOrm()
	err := o.Read(model)
	return err
}

//get article view count
func (model *Article) GetArticleViewCount(id int) (int, error) {
	o := orm.NewOrm()
	count := 0
	err := o.Raw("SELECT count(1) FROM article_log WHERE article_id = ?", id).QueryRow(&count)
	if err != nil {
		return 0, err
	}
	return count, nil

}

//get one article detail data
func (model *Article) ArticleDetail() (ArticlePortal, error) {
	o := orm.NewOrm()
	err := o.Read(model)
	articlePortal := ArticlePortal{}
	strID := strconv.Itoa(int(model.Id))
	//get article view count
	o.Raw("SELECT count(1) FROM article_log WHERE article_id = ?", model.Id).QueryRow(&articlePortal.ViewCount)
	//get article previous
	o.Raw("SELECT id, title FROM article WHERE id < " + strID + "ORDER BY id DESC LIMIT 1;").QueryRow(&articlePortal.Previous)
	//get article next
	o.Raw("SELECT id , title FROM article WHERE id > " + strID + "ORDER BY id ASC LIMIT 1;").QueryRow(&articlePortal.Next)
	return articlePortal, err
}
