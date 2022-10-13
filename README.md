# iqdbAPI
A Go package for anime illustration searching on iqdb.org

## Usage:
```go
f := iqdbAPI.API(iqdbAPI.Options{
Url: "https://xxx.xxx/xxx/xxx.jpg",
FilePath: "X:\\XX\\xxx.jpg"})
// return in JSON
```

## Example:
```go
package main

import (
	"fmt"
	"iqdbAPI"
)


func main() {
	// Url: Search image using an external link
	// FilePath: Search using your local image file (Absolute Path)
	// Url and FilePath are optional parameters, you can use either of them.
	f := iqdbAPI.API(iqdbAPI.Options{
		Url: "https://i.pximg.net/img-master/img/2021/08/19/19/47/13/92104378_p0_master1200.jpg",
		FilePath: "C:\\Users\\Harry\\GolandProjects\\test\\test.jpg"})
	
	
	//Or you can call API1(url) or API2(filepath) directly
	m := iqdbAPI.API2("C:\\Users\\Harry\\GolandProjects\\test\\test.jpg")
	
	q := iqdbAPI.API1("https://i.pximg.net/img-master/img/2021/08/19/19/47/13/92104378_p0_master1200.jpg")
	
	fmt.Println(q)
	
	
	
}
```
```go
// result
{"Success":true,"Results":[{"Head":"Best match","Url":"https://yande.re/post/show/834798","Titles":"Rating: s Score: 2 Tags: kimono morino_rinze sky_cappuccino the_idolm@ster the_idolm@ster_shiny_colors umbrella","Height":"2192","Width":"2824","Category":"[Safe]","Similarity":"98% similarity"},{"Head":"Additional match","Url":"https://yande.re/post/show/787237","Titles":"Rating: s Score: 64 Tags: kimono morino_rinze sky_cappuccino the_idolm@ster the_idolm@ster_shiny_colors umbrella","Height":"2192","Width":"2824","Category":"[Safe]","Similarity":"98% similarity"},{"Head":"Additional match","Url":"//danbooru.donmai.us/posts/4528058","Titles":"Rating: s Tags: 1girl absurdres black_hair eyebrows_visible_through_hair flower forest hair_bun hair_flower hair_ornament highres idolmaster idolmaster_shiny_colors japanese_clothes kimono looking_at_viewer morino_rinze nature outdoors purple_hair red_eyes sidelocks sitting sky_cappuccino solo tied_hair tree umbrella","Height":"2192","Width":"2824","Category":"[Safe]","Similarity":"98% similarity"},{"Head":"Additional match","Url":"//anime-pictures.net/pictures/view_post/719427?lang=en","Titles":"","Height":"2192","Width":"2824","Category":"[Safe]","Similarity":"98% similarity"},{"Head":"Additional match","Url":"//danbooru.donmai.us/posts/4715098","Titles":"Rating: s Score: 1 Tags: 1girl absurdres black_hair commentary eyebrows_visible_through_hair flower forest hair_bun hair_flower hair_ornament highres idolmaster idolmaster_shiny_colors japanese_clothes kimono looking_at_viewer morino_rinze nature outdoors purple_hair red_eyes sidelocks sitting sky_cappuccino solo tied_hair tree umbrella","Height":"2192","Width":"2824","Category":"[Safe]","Similarity":"97% similarity"}]}
```

## Credits
[PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
