package files

import "read-adviser-bot/storage"

type Storage struct {
	basePath string
}

t defaultPerm = 0774
var ErrNoSavedPages = errors.New("no saved pages")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {

	defer func () { err e.WrapIfErr("can't save page", err) } ()

	fPath := filepath.Join(s.basePath, page.UserName)

	//make dir

	if err := os.MkdirAll(path,defaultPerm);err!=nil{
		return err
	}
	
	//forming filename
	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join (filePath, fName) 
	
	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() { _=file.Close()}()

	//gob TODO:learn more

	if err := gob.NewEngoder(file).Encode(page); err!=nil{
		return err
	}

	return nil
}	


func (s Storage) PickRandom(username string) (page *storage.Page) {
	defer func() { err = e.WrapIfErr ("can't pick random page" ,err) }() 
	
	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	// 0-9
	
	


}


func fileName(p *storage.Page) (string, error) {
	return p.Hash()

}
