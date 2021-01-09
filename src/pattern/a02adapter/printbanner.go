package a02adapter

type BannerAdapterStu struct {
	Banner BannerStu
}

func (r *BannerAdapterStu) Init(banner BannerStu) {
	r.Banner = banner

}
func (r *BannerAdapterStu) PrintWeak() {
	r.Banner.ShowWithParen()
}

func (r *BannerAdapterStu) PrintStrong() {
	r.Banner.ShowWithAster()
}
