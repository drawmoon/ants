package stmt

import (
	"testing"

	"github.com/drawmoon/ants/meta"
)

func TestBuildStatement(t *testing.T) {
	metaSet := createDefaultMetaSet()

	f, _ := metaSet.GetTableField("1")
	field1 := SelectPart{Expr: SelectField{Field: f}}

	s, err := Select(field1).String()
	if err != nil {
		t.Error(err)
	}

	if s != "SELECT * FROM " {
		t.Error("build statement error")
	}
}

func createDefaultMetaSet() *meta.MetaSet {
	stud := &meta.Table{
		Id:     "1",
		Name:   "abc_student",
		Schema: "exp_data",
	}
	stud_id := &meta.TableField{
		Id:      "1",
		Name:    "id",
		TableId: stud.Id,
	}
	stud_name := &meta.TableField{
		Id:      "2",
		Name:    "name",
		TableId: stud.Id,
	}
	stud_age := &meta.TableField{
		Id:      "3",
		Name:    "age",
		TableId: stud.Id,
	}
	stud_count := &meta.TableField{
		Id:      "4",
		Name:    "count",
		Exprstr: "COUNT({1})",
		TableId: stud.Id,
	}
	stud_cid := &meta.TableField{
		Id:      "5",
		Name:    "class_id",
		TableId: stud.Id,
	}
	stud_cast := &meta.TableField{
		Id:      "6",
		Name:    "age_cast",
		Exprstr: "CAST({3} AS INT)",
		TableId: stud.Id,
	}
	stud_sex := &meta.TableField{
		Id:      "7",
		Name:    "sex",
		Exprstr: "CASE :self WHEN 0 THEN 'man' ELSE 'woman' END",
		TableId: stud.Id,
	}
	stud_date := &meta.TableField{
		Id:      "8",
		Name:    "enroll_date",
		TableId: stud.Id,
	}
	stud_year := &meta.TableField{
		Id:      "9",
		Name:    "enroll_year",
		Exprstr: "RAW_SQL('substr({8}, 1 ,4)')",
		TableId: stud.Id,
	}

	score := &meta.Table{
		Id:     "2",
		Name:   "abc_score",
		Schema: "exp_data",
	}
	score_id := &meta.TableField{
		Id:      "11",
		Name:    "id",
		TableId: score.Id,
	}
	score_sid := &meta.TableField{
		Id:      "12",
		Name:    "student_id",
		TableId: score.Id,
	}
	score_subj := &meta.TableField{
		Id:      "13",
		Name:    "subject",
		TableId: score.Id,
	}
	score_score := &meta.TableField{
		Id:      "14",
		Name:    "score",
		TableId: score.Id,
	}
	score_subjc := &meta.TableField{
		Id:      "15",
		Name:    "subject_code",
		TableId: score.Id,
	}

	class := &meta.Table{
		Id:     "3",
		Name:   "abc_class",
		Schema: "exp_data",
	}
	class_id := &meta.TableField{
		Id:      "21",
		Name:    "id",
		TableId: class.Id,
	}
	class_name := &meta.TableField{
		Id:      "22",
		Name:    "name",
		TableId: class.Id,
	}

	j1 := &meta.Relationship{
		Id:             "1",
		JoinType:       meta.Left,
		ConditionLeft:  "1",
		ConditionRight: "12",
	}
	j2 := &meta.Relationship{
		Id:             "2",
		JoinType:       meta.Left,
		ConditionLeft:  "12",
		ConditionRight: "1",
	}
	j3 := &meta.Relationship{
		Id:             "3",
		JoinType:       meta.Left,
		ConditionLeft:  "5",
		ConditionRight: "21",
	}
	j4 := &meta.Relationship{
		Id:             "4",
		JoinType:       meta.Left,
		ConditionLeft:  "21",
		ConditionRight: "5",
	}

	t := []*meta.Table{stud, score, class}
	f := []*meta.TableField{
		stud_id, stud_name, stud_age, stud_count, stud_cid, stud_cast, stud_sex, stud_date, stud_year,
		score_id, score_sid, score_subj, score_score, score_subjc, class_id, class_name,
	}
	r := []*meta.Relationship{j1, j2, j3, j4}

	m, err := meta.NewMetaSet("1", t, f, r)
	if err != nil {
		panic(err)
	}

	return m
}
