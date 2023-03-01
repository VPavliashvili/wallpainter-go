package listimages

type WriterNilError struct {}
func (err WriterNilError) Error() string {
    return "writer is nil"
}
