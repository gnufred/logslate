package threat

import "testing"

var testPatterns = []ptrn{
	{Name: "x", Type: Unk, Score: 1},
	{Name: "xy", Type: Unk, Score: 2},
	{Name: "xyz", Type: Unk, Score: 3},

	{Name: "\x99()", Type: Unk, Score: 255},
	{Name: "'SLEEP(0)", Type: SQL, Score: 255},
	{Name: "socket.hitmycommandcenter", Type: CMD, Score: 255},
	{Name: "../etc/passwd", Type: Dir, Score: 255},
	{Name: "<script>", Type: XSS, Score: 255},
}

// TestFalsePositives tests squashed false positives after 2023-07-06. This will avoid
// detecting them again when adding new attack patterns.
// IMPORTANT: Listed queries need to be "anonymized" as best as possible.
func TestFalsePositives(t *testing.T) {
	detectionShouldNotTrigger := func(t *testing.T, id, query, path string) {
		atkType, atkScore := attackScan(path, query)
		if atkType != 0 {
			t.Errorf(
				"%s: false positive found, attak type != 0, got %d",
				id, atkType,
			)
		}
		if atkScore != 0 {
			t.Errorf(
				"%s: false positive found, attak score != 0, got %d",
				id, atkScore,
			)
		}
	}

	falsePositives := []struct {
		// The ID is used in failure output to make it easier to identify what failed.
		// Format is: wX (week X) . int increment.
		Id    string
		Path  string
		Query string
	}{
		{"w1.01", "/graphql", "query={products(currentPage%3A1%2CpageSize%3A60%2Cfilter%3A%7Bcategory_id%3A%7Beq%3A%2217076%22%7D%2Cdata_source%3A%7Bin%3A%5B%22%22%5D%7D%7D%2Csearch%3A%22l%27%C3%A9ducation%20du%20miracle%22%2CresolverLight%3A1){total_count,global_expanded_facets,new_search_result_count,display_message,aggregations{options{count,label,value},attribute_code,count,label,is_swatch,expanded_mode},page_info{category{label,value},color_affiliation{label,value},seller_name%20{label,value},page_size,current_page,total_pages,__typename},items%20{__typename,id,sku,name,erp_product_code,image{name,url},small_image{url},kit_nb_options,kit_image{url},url_key,type_id,ebookType,cross_format,cross_format_label,ean,rating,nb_reviews,front_subtitle,automatic_flags,mp_info{offers{offer_id,ranking,price,total_price,state_code,origin_price},total_new,total_used,lowest_price_new{value},lowest_price_used{value},second_lowest_price_new{value},second_lowest_price_used{value}},manual_flags,main_flag%20{label,value},gsm_salesrule_ids{gsm_salesrule_id,gsm_salesrule_title,gsm_salesrule_banner_title,gsm_salesrule_description,gsm_salesrule_picto,gsm_salesrule_url,gsm_salesrule_offer},price_range{minimum_price{regular_price%20{value,currency},final_price%20{value,currency},discount%20{amount_off,percent_off},final_price_excl_tax%20{value,currency}},maximum_price%20{regular_price%20{value,currency},final_price%20{value,currency},discount%20{amount_off,percent_off}}},...%20on%20ConfigurableProduct{price_range{maximum_price{regular_price%20{value,currency},final_price%20{value,currency},discount%20{amount_off,percent_off}}}},attribute_set_id,attribute_set_name,release_date,backorder_end_date,web_sellable,stock_item_extra{front_availability,availability_date,order_delay,min_quantity_in_cart,offer{availability,seller_code,qty}},book_support,support_musical_custom_:%20support_musical,support_video_custom_:%20support_video,format_book_custom_:%20format_book,button_id_custom_:%20button_id,special_price_custom_:%20special_price}}}"},
	}

	for _, d := range falsePositives {
		detectionShouldNotTrigger(t, d.Id, d.Path, d.Query)
	}
}

func TestAttackTypeString(t *testing.T) {
	testMatch := func(t *testing.T, in AttackType, exp string) {
		got := String(in)
		if got != exp {
			t.Errorf("attack type string doesn't match, got %s, exp %s", got, exp)
		}
	}

	data := []struct {
		In  AttackType
		Exp string
	}{
		{Nil, ""},
		{Unk, "generic"},
		{SQL, "sql_injection"},
		{CMD, "command_injection"},
		{Dir, "directory_traversal"},
		{XSS, "xss"},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}
