package dao

import (
	"github.com/mitchellh/mapstructure"
	"github.com/sillyhatxu/mysql-client"
	"word-list/dto"
)

const (
	page_sql = `
		select id,word,context
		from word_list
		where word_group_id = ?
		order by id asc 
		limit ?,?
	`
)

func FindWordListPageByParams(id int64, currentPage int64, pageSize int64) ([]dto.WordList, error) {
	limit := currentPage*pageSize + 1
	results, err := dbclient.Client.Find(page_sql, id, limit, pageSize)
	if err != nil {
		return nil, err
	}
	var wordArray []dto.WordList
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		WeaklyTypedInput: true,
		Result:           &wordArray,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(results)
	if err != nil {
		return nil, err
	}
	return wordArray, nil
}
