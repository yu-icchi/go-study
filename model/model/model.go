package model

import (
	"./admin"
	"./user"
	"./family"
)

var (
	User = &user.User{}
	UserFile = &user.File{}
	Family = &family.Family{}
	FamilyAvatar = &family.Avatar{}
	FamilySticker = &family.Sticker{}
	FamilyParts = &family.Parts{}
	FamilyColor = &family.Color{}
	FamilyColorSet = &family.ColorSet{}
	AdminUser = &admin.User{}
	AdminRole = &admin.Role{}
)

func Collections() []string {
	return []string{
		User.Collection(),
		UserFile.Collection(),
		Family.Collection(),
		FamilyAvatar.Collection(),
		FamilySticker.Collection(),
		FamilyParts.Collection(),
		FamilyColor.Collection(),
		FamilyColorSet.Collection(),
		AdminUser.Collection(),
		AdminRole.Collection(),
	}
}
