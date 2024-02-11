package types

type School struct{
	Name string
	NumberOfStudents int

}

func (s *School) UpDateSchool(name string,number int){
	s.Name =name
	s.NumberOfStudents = number
}

func (s School) GetSchoolDetails()(string,int){
	return s.Name, s.NumberOfStudents
}