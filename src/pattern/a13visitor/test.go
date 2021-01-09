package a13visitor

func TestFunc() {
	rootdir := &DirectoryStu{Name: "root"}
	bindir := &DirectoryStu{Name: "bin"}
	tmpdir := &DirectoryStu{Name: "tmp"}
	usrdir := &DirectoryStu{Name: "usr"}
	rootdir.AddEntry(bindir)
	rootdir.AddEntry(tmpdir)
	rootdir.AddEntry(usrdir)

	bindir.AddEntry(&FileStu{Name: "vi.doc", Size: 10000})
	bindir.AddEntry(&FileStu{Name: "latex.txt", Size: 20000})
	var VisitorList VisitorListStu
	VisitorList.Reset("&&")
	rootdir.Accept(&VisitorList)
}
