package surreal

func Delete(id string) error {
	_, err := DB.Delete(id)
	return err
}
