package checkerr

func Checkerr(err error){
	if err != nil{
		panic(err)
	}
}