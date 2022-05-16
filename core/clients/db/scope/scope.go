package scope

import "gorm.io/gorm"

func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = ?", true)
}

func IsDeletedT(tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".is_deleted = ?", true)
	}
}

func IsNotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = ?", false)
}

func IsNotDeletedT(tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".is_deleted = ?", false)
	}
}

func IsActive(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", true)
}

func IsActiveT(tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".is_active = ?", true)
	}
}

func IsNotActive(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", false)
}

func IsNotActiveT(tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".is_active = ?", false)
	}
}

func IsExpired(db *gorm.DB) *gorm.DB {
	return db.Where("is_expired = ?", true)
}

func IsExpiredT(tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".is_expired = ?", true)
	}
}

func IsNotExpired(db *gorm.DB) *gorm.DB {
	return db.Where("is_expired = ?", false)
}

func IsNotExpiredT(tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".is_expired = ?", false)
	}
}
