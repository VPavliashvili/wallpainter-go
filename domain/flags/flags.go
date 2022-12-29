package flags

const (
	Help         = "--help"
	SetWallpaper = "--set-wallpaper"
)

type Flag string

//there was an attemp of enum implementation
//and I ended up conflicting with unit test
//eventually these two functions stays since
//it will take lots of editing to remove these two
func (f Flag) String() string {
    return string(f)
}

func ToFlag(s string) Flag {
    return Flag(s)
}
