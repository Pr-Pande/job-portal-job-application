package database

import "job-portal-api/internal/models"

func (r *Repo) AutoMigrate() error {
	//if s.db.Migrator().HasTable(&User{}) {
	//	return nil
	//}

	
	err := r.DB.Migrator().DropTable()
	if err != nil {
		return err
	}
	
	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = r.DB.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}

	return nil
}
