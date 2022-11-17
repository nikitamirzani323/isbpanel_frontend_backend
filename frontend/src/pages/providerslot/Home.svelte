<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "PROVIDER SLOT"
    let sData = "";
    let myModal_newentry = "";

    let idrecord = "";
    let name_field = "";
    let display_field = 0;
    let slug_field = "";
    let image_field = "";
    let title_field = "";
    let descp_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";

    let listgame = [];
    let provider_name = "";
    let provider_slug = "";

    let sDataPrediksi = "";
    let prediksislot_id_field = 0
    let providerslot_id_field = 0
    let prediksislot_prediksi_field = 0
    let prediksislot_name_field = ""
    let prediksislot_image_field = ""
    let prediksislot_status_field = ""
    let prediksislot_create_field = ""
    let prediksislot_update_field = ""

  
    let searchProviderSlot = "";
    let filterProviderSlot = [];
    let searchGameSlot = "";
    let filterGameSlot = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchProviderSlot) {
            filterProviderSlot = listHome.filter(
                (item) =>
                    item.domain_name
                        .toLowerCase()
                        .includes(searchProviderSlot.toLowerCase())
            );
        } else {
            filterProviderSlot = [...listHome];
        }
        if (searchGameSlot) {
            filterGameSlot = listgame.filter(
                (item) =>
                    item.prediksislot_name
                        .toLowerCase()
                        .includes(searchGameSlot.toLowerCase())
            );
        } else {
            filterGameSlot = [...listgame];
        }
    }
    const NewData = (e,id,name,display,image,slug,title,descp,status,createdata,updatedata) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            idrecord = parseInt(id)
            name_field = name;
            display_field = display;
            slug_field = slug;
            image_field = image;
            title_field = title;
            descp_field = descp;
            status_field = status;
            create_field = createdata;
            update_field = updatedata;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const NewGameSlot = (id,name,slug) => {
        clearField()
        call_listgameslot(id)
        providerslot_id_field = id
        provider_name = name
        provider_slug = slug
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistgame"));
        myModal_newentry.show();
    };
    const ShowNewGame = (e,idgame,prediksigame,nmgame,imggame,statusgame,creategame,updategame) => {
        sDataPrediksi = e
        if(e == "Edit"){
            prediksislot_id_field = idgame
            prediksislot_prediksi_field = prediksigame
            prediksislot_name_field = nmgame
            prediksislot_image_field = imggame
            prediksislot_status_field = statusgame
            prediksislot_create_field = creategame
            prediksislot_update_field = updategame
        }else{
            clearFieldGame();
        }
        myModal_newentry = new bootstrap.Modal(
            document.getElementById("modalentrygame")
        );
        myModal_newentry.show();
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "Edit"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
        }
        if(name_field == ""){
            flag = false
            msg += "The Name is required\n"
        }
        if(slug_field == ""){
            flag = false
            msg += "The Slug is required\n"
        }
        if(image_field == ""){
            flag = false
            msg += "The Image is required\n"
        }
        if(title_field == ""){
            flag = false
            msg += "The Title SEO is required\n"
        }
        if(descp_field == ""){
            flag = false
            msg += "The Descp SEO is required\n"
        }
        if(status_field == ""){
            flag = false
            msg += "The Status is required\n"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/providerslotsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"SLOT-SAVE",
                    providerslot_id: parseInt(idrecord),
                    providerslot_name: name_field,
                    providerslot_display: parseInt(display_field),
                    providerslot_image: image_field,
                    providerslot_slug: slug_field,
                    providerslot_title: title_field,
                    providerslot_descp: descp_field,
                    providerslot_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSaveGameSlot() {
        let flag = true
        let msg = ""
        if(sDataPrediksi == "New"){
            if(providerslot_id_field == ""){
                flag = false
                msg += "The Provider is required\n"
            }
            if(prediksislot_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(prediksislot_image_field == ""){
                flag = false
                msg += "The Image is required\n"
            }
            if(prediksislot_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
            if(prediksislot_prediksi_field < 1 || prediksislot_prediksi_field > 99){
                flag = false
                msg += "The Range Prediksi 0 - 99\n"
            }
        }else{
            if(prediksislot_id_field == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(providerslot_id_field == ""){
                flag = false
                msg += "The Provider is required\n"
            }
            if(prediksislot_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(prediksislot_image_field == ""){
                flag = false
                msg += "The Image is required\n"
            }
            if(prediksislot_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
            if(prediksislot_prediksi_field < 1 || prediksislot_prediksi_field > 99){
                flag = false
                msg += "The Range Prediksi 0 - 99\n"
            }
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/prediksislotsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataPrediksi,
                    page:"DOMAIN-SAVE",
                    prediksislot_id: parseInt(prediksislot_id_field),
                    providerslot_id: parseInt(providerslot_id_field),
                    prediksislot_prediksi: parseInt(prediksislot_prediksi_field),
                    providerslot_slug: provider_slug,
                    prediksislot_name: prediksislot_name_field,
                    prediksislot_image: prediksislot_image_field,
                    prediksislot_status: prediksislot_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                if(sDataPrediksi=="New"){
                    clearFieldGame()
                }
                call_listgameslot(providerslot_id_field)
                RefreshHalaman()
                msgloader = json.message;
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleDeleteGameSlot(idprediksislot) {
        let flag = true;
        let msg = "";
        if (idprediksislot == "" || parseInt(idprediksislot) == 0) {
            flag = false;
            msg = "The Id Game is required";
        }
        if (providerslot_id_field == "" || parseInt(providerslot_id_field) == 0) {
            flag = false;
            msg = "The Id Provider is required";
        }
        if (flag) {
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/prediksislotdelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "SLOT-SAVE",
                    prediksislot_id: parseInt(idprediksislot),
                    providerslot_id: parseInt(providerslot_id_field),
                    providerslot_slug: provider_slug,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                call_listgameslot(providerslot_id_field)
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function handleGeneratorSlot(idprovider,slug) {
        let flag = true;
        let msg = "";
        if (idprovider == "" || parseInt(idprovider) == 0) {
            flag = false;
            msg = "The Id Provider is required";
        }
        if (flag) {
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/prediksislotgenerator", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "SLOT-SAVE",
                    providerslot_id: parseInt(idprovider),
                    providerslot_slug: slug,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                RefreshHalaman()
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function call_listgameslot(e) {
        listgame = [];
        css_loader = "display: inline-block;";
        msgloader = "Fetching...";
        const res = await fetch("/api/prediksislot", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "KELUARAN-view",
                providerslot_id: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let prediksi_statuscss = "";
                let prediksi_css = "";
                let prediksi_css_color = "";
                let prediksi_class = "";
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["prediksislot_status"] == "Y"){
                        prediksi_statuscss = "background:#FFEB3B;font-weight:bold;color:black;"
                    }else{
                        prediksi_statuscss = "background:#E91E63;font-weight:bold;color:white;"
                    }
                    if(parseInt(record[i]["prediksislot_prediksi"]) > 9 && parseInt(record[i]["prediksislot_prediksi"]) < 50){
                        prediksi_css = "color: #dc3545;"
                        prediksi_css_color = "color: white;"
                        prediksi_class = "bg-danger"
                    }
                    if(parseInt(record[i]["prediksislot_prediksi"]) > 49 && parseInt(record[i]["prediksislot_prediksi"]) < 71){
                        prediksi_css = "color: #ffc107;"
                        prediksi_css_color = "color: black;"
                        prediksi_class = "bg-warning"
                    }
                    if(parseInt(record[i]["prediksislot_prediksi"]) > 70 && parseInt(record[i]["prediksislot_prediksi"]) < 101){
                        prediksi_css = "color: #28a745;"
                        prediksi_css_color = "color: white;"
                        prediksi_class = "bg-success"
                    }
                    listgame = [
                        ...listgame,
                        {
                            prediksislot_id: record[i]["prediksislot_id"],
                            prediksislot_nmprovider: record[i]["prediksislot_nmprovider"],
                            prediksislot_name: record[i]["prediksislot_name"],
                            prediksislot_prediksi: record[i]["prediksislot_prediksi"],
                            prediksislot_prediksicss: prediksi_css,
                            prediksislot_prediksicss_color: prediksi_css_color,
                            prediksislot_prediksiclass: prediksi_class,
                            prediksislot_image: record[i]["prediksislot_image"],
                            prediksislot_status: record[i]["prediksislot_status"],
                            prediksislot_statuscss: prediksi_statuscss,
                            prediksislot_create: record[i]["prediksislot_create"],
                            prediksislot_update: record[i]["prediksislot_update"],
                        },
                    ];
                }
            }
        }
        msgloader = json.message;
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    function clearField(){
        idrecord = "";
        name_field = "";
        display_field = 0;
        slug_field = "";
        image_field = "";
        title_field = "";
        descp_field = "";
        status_field = "";
        create_field = "";
        update_field = "";

        prediksislot_id_field = 0
        providerslot_id_field = 0
        prediksislot_prediksi_field = 0
        prediksislot_name_field = ""
        prediksislot_image_field = ""
        prediksislot_status_field = ""
    }
    function clearFieldGame(){
        prediksislot_id_field = 0
        prediksislot_prediksi_field = 0
        prediksislot_name_field = ""
        prediksislot_image_field = ""
        prediksislot_status_field = ""
        prediksislot_create_field = ""
        prediksislot_update_field = ""
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
            case "NEW_GAME":
                ShowNewGame("New",0,0,"","","");
            break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
    const handleKeyboard_format = () => {
        let numbera;
        for (let i = 0; i < display_field.length; i++) {
            numbera = parseInt(display_field[i]);
            if (isNaN(numbera)) {
                display_field = "";
            }
        }
        for (let i = 0; i < prediksislot_prediksi_field.length; i++) {
            numbera = parseInt(prediksislot_prediksi_field[i]);
            if (isNaN(numbera)) {
                prediksislot_prediksi_field = "";
            }
        }
    };
    function status(e){
        let result = "HIDE"
        if(e == "Y"){
            result = "SHOW"
        }
        return result
    }
    function lowercase(element) {
        function onInput(event) {
            element.value = element.value.toLowerCase();
        }
        element.addEventListener("input", onInput);
        return {
            destroy() {
                element.removeEventListener("input", onInput);
            },
        };
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchProviderSlot}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Provider Slot"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="3">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">IMAGE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PROVIDER</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">SLUG</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTAL GAME</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">COUNTER</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DISPLAY</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterProviderSlot as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewData("Edit",
                                                rec.providerslot_id,rec.providerslot_name, rec.providerslot_display,
                                                rec.providerslot_image,rec.providerslot_slug, rec.providerslot_title,
                                                rec.providerslot_descp,rec.providerslot_status, rec.providerslot_create,rec.providerslot_update);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewGameSlot(rec.providerslot_id,rec.providerslot_name,rec.providerslot_slug);
                                            }} 
                                            class="bi bi-file-earmark-plus"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                handleGeneratorSlot(rec.providerslot_id,rec.providerslot_slug);
                                            }} 
                                            class="bi bi-gear"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.providerslot_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.providerslot_statuscss}">
                                            {status(rec.providerslot_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <img src="{rec.providerslot_image}" width="100" alt="">
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.providerslot_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.providerslot_slug}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font}; {rec.providerslot_totalgameslot_css}">{rec.providerslot_totalgameslot}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font}; {rec.providerslot_counter_css}">{rec.providerslot_counter}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font}; {rec.providerslot_display_css}">{rec.providerslot_display}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input
                        bind:value={name_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Slug</label>
                    <input
                      use:lowercase
                      bind:value={slug_field}
                      class="required form-control"
                      type="text"
                      placeholder="Slug"/>
                    <p class="text-muted">Ex : pragmatic-play or slot-joker</p>
                  </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Link Image</label>
                    <Input
                        bind:value={image_field}
                        class="required"
                        type="text"
                        placeholder="Link Image"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Display</label>
                    <Input
                      bind:value={display_field}
                      on:keyup={handleKeyboard_format}
                      class="required"
                      type="text"
                      style="text-align:right;"
                      placeholder="Display"/>
                  </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select class="form-control required" bind:value="{status_field}">
                        <option value="Y">SHOW</option>
                        <option value="N">HIDE</option>
                    </select>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Title - Seo</label>
                    <Input
                        bind:value={title_field}
                        class="required"
                        type="text"
                        maxlength=60
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Deskripsi - Seo</label>
                    <textarea
                      style="height: 100px;resize: none;"
                      bind:value={descp_field}
                      class="form-control required"/>
                </div>
                <div class="alert alert-info" role="alert" style="font-size:12px;padding:5px;">
                    Ex :<br /> 
                    Title SEO : <br />Slot Gacor Hari ini | RTP Slot Spade Gaming
                    <br /> <br />
                    Descp SEO : <br />RTP Slot Gacor Hari ini | Bocoran RTP Slot Gacor | RTP Slot Spade Gaming | Bocoran Slot Gacor Hari ini | RTP Slot Gampang Menang | Situs Slot Gacor Hari ini
                </div>
                {#if sData == "Edit"}
                    <div class="alert alert-dark" role="alert" style="font-size:11px;padding:5px;">
                        Create : {create_field}
                        <br />
                        Update : {update_field}
                    </div>
                {/if}
            </div>
        </div>
        
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>

<Modal
  modal_id="modallistgame"
  modal_size="modal-dialog-centered modal-lg"
  modal_title="Provider : {provider_name}"
  modal_body_css="padding:5px;margin:0px;height:500px;overflow-y: scroll; "
  modal_footer_css="padding:5px;"
  modal_search={true}
  modal_footer={true}>
  <slot:template slot="search">
    <input
        bind:value={searchGameSlot}
        type="text"
        class="form-control"
        placeholder="Search Game Slot"
        aria-label="Search"/>
  </slot:template>
  <slot:template slot="body">
    <table class="table">
      <thead>
        <tr>
          <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};" colspan=2>&nbsp;</th>
          <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
          <th NOWRAP width="20%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">GAME</th>
          <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">IMAGE</th>
          <th NOWRAP width="*" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PREDIKSI</th>
        </tr>
      </thead>
      <tbody>
        {#each filterGameSlot as rec}
          <tr>
            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
              <i
                on:click={() => {
                  handleDeleteGameSlot(rec.prediksislot_id,rec.prediksislot_prediksi);
                }}
                class="bi bi-trash"/>
            </td>
            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                <i
                  on:click={() => {
                    ShowNewGame("Edit",rec.prediksislot_id,rec.prediksislot_prediksi,rec.prediksislot_name,rec.prediksislot_image,rec.prediksislot_status,rec.prediksislot_create,rec.prediksislot_update);
                  }}
                  class="bi bi-pencil"/>
              </td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.prediksislot_statuscss}">{status(rec.prediksislot_status)}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.prediksislot_name}</td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                <img src="{rec.prediksislot_image}" width="40" alt="">
            </td>
            <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">
                <div class="progress">
                    <div class="progress-bar progress-bar-striped {rec.prediksislot_prediksiclass}" role="progressbar" style="width: {rec.prediksislot_prediksi}%" aria-valuenow="{rec.prediksislot_prediksi}" aria-valuemin="0" aria-valuemax="100">
                        <span style="{rec.prediksislot_prediksicss_color}">{rec.prediksislot_prediksi}%</span>
                    </div>
                </div>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="NEW_GAME"
      button_title="New"
      button_css="btn-primary"/>
  </slot:template>
</Modal>

<Modal
	modal_id="modalentrygame"
	modal_size="modal-dialog-centered "
	modal_title="{sDataPrediksi} Game"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <Input
                bind:value={prediksislot_name_field}
                class="required"
                type="text"
                placeholder="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Link Image</label>
            <Input
                bind:value={prediksislot_image_field}
                class="required"
                type="text"
                placeholder="Link Image"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Prediksi</label>
            <Input
              bind:value={prediksislot_prediksi_field}
              on:keyup={handleKeyboard_format}
              class="required"
              type="text"
              maxlength="2"
              style="text-align:right;"
              placeholder="Prediksi"/>
          </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select class="form-control required" bind:value="{prediksislot_status_field}">
                <option value="Y">SHOW</option>
                <option value="N">HIDE</option>
            </select>
        </div>
        {#if sDataPrediksi == "Edit"}
            <div class="alert alert-dark" role="alert" style="font-size:11px;padding:5px;">
                Create : {prediksislot_create_field}
                <br />
                Update : {prediksislot_update_field}
            </div>
        {/if}
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={() => {
                handleSaveGameSlot();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>