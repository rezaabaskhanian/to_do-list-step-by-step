package Assignee


import "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"


//TODO: struct ro maloom kon va interface ro bede beehsh 


// NEW kon azash va address struct jadidi ke mikhad sakhte beshe ro bargardoon 


// function haye in struct ro map kon behesh va harja niaz dari repository ro bekhoni azash bekhoon mesle save va load ke dar task estefade kardi 



// CreateAssignee یک Assignee جدید ایجاد می‌کند
func CreateAssignee(name, email string) domain.Assignee {
	// به طور موقت ID را به صورت دستی می‌دهیم
	assignee := domain.Assignee{
		ID:    1, // در آینده باید این ID از دیتابیس یا حافظه گرفته بشه
		Name:  name,
		Email: email,
	}
	return assignee
}
