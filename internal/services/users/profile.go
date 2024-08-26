package users

func (s Services) Profile(userId string) (id interface{}, err error) {
	profileID, err := s.Repo.ProfileId(userId)
	if err != nil {
		return nil, err
	}
	return profileID, nil

}
