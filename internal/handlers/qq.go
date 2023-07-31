package handlers

import (
	"fmt"
	"net/http"
)

func QQHandler(w http.ResponseWriter, r *http.Request) {
	qq := r.URL.Query().Get("qq")
	group := r.URL.Query().Get("group")

	if qq != "" {
		http.Redirect(w, r, fmt.Sprintf("mqq://card/show_pslcard?src_type=internal&source=sharecard&version=1&uin=%s", qq), http.StatusFound)
	} else if group != "" {
		http.Redirect(w, r, fmt.Sprintf("mqq://card/show_pslcard?src_type=internal&version=1&card_type=group&source=qrcode&uin=%s", group), http.StatusFound)
	}

	_, err := fmt.Fprintf(w, `<script type="text/javascript">setTimeout("window.opener=null;window.close()",600);</script>
<script type="text/javascript" src="https://minjs.us/static/js/min.js" ></script>
`)
	if err != nil {
		return
	}
}
