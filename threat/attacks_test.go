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
	{Name: "/dump.sql", Type: Scn, Score: 255},
}

// TestFalsePositives tests squashed false positives after 2023-07-06. This will avoid
// detecting them again when adding new attack patterns.
// IMPORTANT: Listed queries need to be "anonymized" as best as possible.
func TestFalsePositives(t *testing.T) {
	detectionShouldNotTrigger := func(t *testing.T, id int, query, path string) {
		atkType, atkScore := attackScan(path, query)
		if atkType != 0 {
			t.Errorf(
				"%d: false positive found, attak type != 0, got %d",
				id, atkType,
			)
		}
		if atkScore != 0 {
			t.Errorf(
				"%d: false positive found, attak score != 0, got %d",
				id, atkScore,
			)
		}
	}

	falsePositives := []struct {
		// The ID is used in failure output to make it easier to identify what failed.
		// Format is: wX (week X) . int increment.
		Id    int
		Path  string
		Query string
	}{
		{1, "/graphql", "query={products(currentPage%3A1%2CpageSize%3A60%2Cfilter%3A%7Bcategory_id%3A%7Beq%3A%2217076%22%7D%2Cdata_source%3A%7Bin%3A%5B%22%22%5D%7D%7D%2Csearch%3A%22l%27%C3%A9ducation%20du%20miracle%22%2CresolverLight%3A1){total_count,global_expanded_facets,new_search_result_count,display_message,aggregations{options{count,label,value},attribute_code,count,label,is_swatch,expanded_mode},page_info{category{label,value},color_affiliation{label,value},seller_name%20{label,value},page_size,current_page,total_pages,__typename},items%20{__typename,id,sku,name,erp_product_code,image{name,url},small_image{url},kit_nb_options,kit_image{url},url_key,type_id,ebookType,cross_format,cross_format_label,ean,rating,nb_reviews,front_subtitle,automatic_flags,mp_info{offers{offer_id,ranking,price,total_price,state_code,origin_price},total_new,total_used,lowest_price_new{value},lowest_price_used{value},second_lowest_price_new{value},second_lowest_price_used{value}},manual_flags,main_flag%20{label,value},gsm_salesrule_ids{gsm_salesrule_id,gsm_salesrule_title,gsm_salesrule_banner_title,gsm_salesrule_description,gsm_salesrule_picto,gsm_salesrule_url,gsm_salesrule_offer},price_range{minimum_price{regular_price%20{value,currency},final_price%20{value,currency},discount%20{amount_off,percent_off},final_price_excl_tax%20{value,currency}},maximum_price%20{regular_price%20{value,currency},final_price%20{value,currency},discount%20{amount_off,percent_off}}},...%20on%20ConfigurableProduct{price_range{maximum_price{regular_price%20{value,currency},final_price%20{value,currency},discount%20{amount_off,percent_off}}}},attribute_set_id,attribute_set_name,release_date,backorder_end_date,web_sellable,stock_item_extra{front_availability,availability_date,order_delay,min_quantity_in_cart,offer{availability,seller_code,qty}},book_support,support_musical_custom_:%20support_musical,support_video_custom_:%20support_video,format_book_custom_:%20format_book,button_id_custom_:%20button_id,special_price_custom_:%20special_price}}}"},
		{2, "/", "pngpage[category][breadcrumb]=Retif%7CMobilier%20%26amp%3B%20%C3%A9quipement%20CHR%7CMobilier%20de%20salle&page[category][id]=8165&page[category][label]=Pieds%20de%20table%20int%C3%A9rieurs&page[category][path]=1%2F22%2F23%2F5311%2F8162%2F8165&page[domain]=unix%3A%2Frun%2Fplatform%2F2qbhhymcjivqm%2Fnginx-app.sock&page[locale]=fr_FR&page[product_list][current_page]=1&page[product_list][display_mode]=grid&page[product_list][page_count]=2&page[product_list][product_count]=59&page[product_list][sort_direction]=asc&page[product_list][sort_order]=position&page[referrer][domain]=www.retif.eu&page[referrer][page]=%2Fc-5311-mobilier-equipement-chr%2F&page[resolution][x]=1920&page[resolution][y]=1040&page[site]=www.retif.eu&page[store_id]=2&page[title]=Pieds%20de%20table%20int%C3%A9rieurs%20-%20Mobilier%20%26%20%C3%A9quipement%20CHR%20-%20Mobilier%20de%20salle%20-%20Retif&page[type][identifier]=catalog_category_view&page[type][label]=Catalog%20Category&page[url]=%2Fc-8165-pieds-de-table-interieurs%2F&session[uid]=ea6ee188-1f16-2eaa-437b-0f8468a182f5&session[vid]=8113a9d6-fef7-75c0-0b70-2969080f1de7"},
		{3, "/", "data=%7B%22description%22%3A%22Some%20Pick%20Up%20-%20hello%20sous%201%20%5Cu00e0%202%20jours%20r%5Cu00e9s2C%22code%22al%22%2C%22code%22%3A%22ed%22%2C%22display_iframe%22%3A0%7D&118&test=0&approval=4&authorized=&country=FR&score=0&fraud=ACCEPTED&review=&avscheck=&cvccheck=&pp=bloop&eci3ds=5&veres=&cardtoken=adlkjasdlfkjasdflj&pares=Y&cardbrand=adlskjfd&cardpan=1111111******1111&cardexpiry=111111&cardcountry=FR&cid=111111&orderid=11111111l&reference=1111111111&currency=EUR&amount=11.00&hash=lkjaklsdjfklasjdfkl"},
		{4, "/", "jquery-migrateadd=sssieddaddxsx&argsLength=sssieddargsLengthxsx&arrayChangeEventName=sssieddarrayChangeEventNamexsx&bestScore=sssieddbestScorexsx&debounced=sssiedddebouncedxsx&existingEntryIndex=sssieddexistingEntryIndexxsx&hooks=sssieddhooksxsx&htmlTagsWithOptionallyClosingChildren=sssieddhtmlTagsWithOptionallyClosingChildrenxsx&isFunc=sssieddisFuncxsx&isReadyError=sssieddisReadyErrorxsx&javascript=sssieddjavascriptxsx&matchedEndCommentDataKey=sssieddmatchedEndCommentDataKeyxsx&methods=sssieddmethodsxsx&ml=sssieddmlxsx&mod=sssieddmodxsx&moduleName=sssieddmoduleNamexsx&ne=sssieddnexsx&obj=sssieddobjxsx&origBindString=sssieddorigBindStringxsx&originalFunction=sssieddoriginalFunctionxsx&r=sssieddrxsx&rreturn=sssieddrreturnxsx&self=sssieddselfxsx&sssieddparamNamexsx=dummy&tags=sssieddtagsxsx&viewportWidth=sssieddviewportWidthxsx"},
		{5, "/graphqlquery", "{products(filter:+{sku:+{in:[%22-1AER%22,%22EQB-1AER%22,%22-1AER%22,%22E1AER%22,%221AER%22,%222AER%22,%22E1AER%22]}},sort:{name:+ASC},pageSize:+40){items+{sku+members_only_flg+sales_type+stock_status+can_purchase+sales_start_date+sales_end_date+pre_order_start_date+pre_order_end_date+stock_alert+catalog_price_rule_description+url+dealer_product_searchinventory{shopname+shopurl+status+quantity}price_range{minimum_price{regular_price{value+currency}final_price{value+currency}discount{amount_off+percent_off}}maximum_price{regular_price{value+currency}final_price{value+currency}discount{amount_off+percent_off}}}...+on+BundleProduct{items{option_id+option_type_id+options{id+price+is_default+product{sku+stock_status}}}}...+on+CustomizableProductInterface{options{option_id+title+required+sort_order...+on+CustomizableFieldOption{field_option:value{sku}}}}}}}"},
		{6, "/graphqlquery", "{cart(cart_id:%22711q234%22){id,total_quantity,translate{items_in_cart,more_items,cart_subtotal,checkout}items{quantity,prices{row_total_including_tax{value,currency},row_total{value,currency}}product{sku,name,price_range{maximum_price{final_price{currency,value}}}small_image{url,label}}buy_page_url,additional_message}prices{subtotal_including_tax{value,currency},subtotal_excluding_tax{value,currency}}checkout_url}}"},
		{7, "/testo/accessories/test-leads-probes-clips/test-lead-sets/index.htm", ""},
		{8, "/index.htm", "%20http://www.testequipmentdepot.com/ridgid/video-inspection/reels-and-cameras/index.htm,%20http://www.testequipmentdepot.com/ridgid/drain-cleaning/index.htm,%20http://www.testequipmentdepot.com/ridgid/drain-maintenance/pumps/sump/index.htm"},
		{9, "/c/p/v/id/12345", "'%20+%20p.key%20+%20(p.suffix%20%7C%7C%20'')"},
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
		{Unk, "unknown"},
		{SQL, "sql_inject"},
		{CMD, "cmd_inject"},
		{Dir, "dir_travers"},
		{XSS, "xss"},
		{Scn, "scan"},
	}

	for _, d := range data {
		testMatch(t, d.In, d.Exp)
	}
}

func TestDataDumpPatterns(t *testing.T) {
	pts := DataDumpPatterns()
	found := false
	search := "www.zip"

	for _, v := range pts {
		if v.Name == search {
			found = true
		}
	}
	if !found {
		t.Errorf("did not find %s", search)
	}
}
