package setwallpaper_test

type cmdFakeArg struct {
	name  string
	value string
}

func (f cmdFakeArg) GetName() string {
	return f.name
}
func (f cmdFakeArg) String() string {
	return f.name
}
func (f cmdFakeArg) Value() string {
	return f.value
}
func (f cmdFakeArg) Description() string {
	return "fake"
}

type fakeio struct{}

func (f fakeio) Exist(file string) bool {
	return file == "validpath.pnj" || file == "validfile"
}

func (f fakeio) IsPicture(file string) bool {
	return file == "validpath.pnj"
}
func (f fakeio) SetWallpaper(file string) error {
	return nil
}

