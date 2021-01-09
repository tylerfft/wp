package a02adapter

func TestFunc() {

	var banner BannerStu
	banner.Init("shi")
	var bannerAdapter BannerAdapterStu
	bannerAdapter.Init(banner)

	PrintFunc(&bannerAdapter)

}

func PrintFunc(prnIF PrintIF) {
	prnIF.PrintWeak()
	prnIF.PrintStrong()
}
