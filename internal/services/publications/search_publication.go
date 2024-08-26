package publications

func (s Services) GetPublicationByID(publiId string) (id interface{}, err error) {
	return s.Repo.FindByID(publiId)
}
