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
	let title_page = "BANNER"
    let sData = "";
    let myModal_newentry = "";
    let domain_field = "";
    let status_field = "";

    let nmbanner_field = "";
    let urlbanner_field = "";
    let urlwebsite_field = "";
    let devicebanner_field = "";
    let posisibanner_field = "";
    let displaybanner_field = 0;
    let statusbanner_field = "";
    let idrecord = "";
    let searchDomain = "";
    let filterDomain = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchDomain) {
            filterDomain = listHome.filter(
                (item) =>
                    item.domain_name
                        .toLowerCase()
                        .includes(searchDomain.toLowerCase())
            );
        } else {
            filterDomain = [...listHome];
        }
    }
    
    const NewData = (e,id,nm,url,urlwebsite,device,posisi,display,status) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            if(sData == "Edit"){
                idrecord = parseInt(id)
            }
            nmbanner_field = nm;
            urlbanner_field = url;
            urlwebsite_field = urlwebsite;
            devicebanner_field = device;
            posisibanner_field = posisi;
            displaybanner_field = display;
            statusbanner_field = status;
        }
        if(sData == "New" || sData == "Edit"){
            myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
            myModal_newentry.show();
        }
        if(sData == "Duplicate"){
                sData = "New"
                idrecord = "";
                handleSave()
        }
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
        if(nmbanner_field == ""){
            flag = false
            msg += "The Banner is required\n"
        }
        if(urlwebsite_field == ""){
            flag = false
            msg += "The URL WEBSITE is required\n"
        }
        if(urlbanner_field == ""){
            flag = false
            msg += "The LINK IMAGE is required\n"
        }
        if(posisibanner_field == ""){
            flag = false
            msg += "The Posisi is required\n"
        }
        if(devicebanner_field == ""){
            flag = false
            msg += "The Device is required\n"
        }
        if(statusbanner_field == ""){
            flag = false
            msg += "The Status is required\n"
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/bannersave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"BANNER-SAVE",
                    banner_id: parseInt(idrecord),
                    banner_name: nmbanner_field,
                    banner_url: urlbanner_field,
                    banner_urlwebsite: urlwebsite_field,
                    banner_posisi: posisibanner_field,
                    banner_device: devicebanner_field,
                    banner_display: parseInt(displaybanner_field),
                    banner_status: statusbanner_field,
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
    function clearField(){
        idrecord = "";
        nmbanner_field = "";
        urlbanner_field = "";
        urlwebsite_field = "";
        devicebanner_field = "";
        posisibanner_field = "";
        displaybanner_field = 0;
        statusbanner_field = "";
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
        for (let i = 0; i < displaybanner_field.length; i++) {
            numbera = parseInt(displaybanner_field[i]);
            if (isNaN(numbera)) {
                displaybanner_field = 0;
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
                            bind:value={searchDomain}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Banner"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan=2>&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DEVICE</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">POSISI</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">BANNER</th>
                                <th NOWRAP width="20%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WEBSITE</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">IMAGE</th>
                                <th NOWRAP width="5%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DISPLAY</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterDomain as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewData("Edit",rec.banner_id,rec.banner_name, rec.banner_url,rec.banner_urlwebsite,rec.banner_device,rec.banner_posisi,rec.banner_display,rec.banner_status);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewData("Duplicate",rec.banner_id,rec.banner_name, rec.banner_url,rec.banner_urlwebsite,rec.banner_device,rec.banner_posisi,rec.banner_display,rec.banner_status);
                                            }} 
                                            class="bi bi-files"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.banner_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.banner_statuscss}">
                                            {status(rec.banner_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.banner_device}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.banner_posisi}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.banner_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <a href="{rec.banner_urlwebsite}" target="_blank">
                                            {rec.banner_urlwebsite}
                                        </a>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <img src="{rec.banner_url}" alt="" width="100">
                                    </td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{rec.banner_display}</td>
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
            <div class="col-lg-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Banner</label>
                    <Input
                        bind:value={nmbanner_field}
                        class="required"
                        type="text"
                        placeholder="Banner"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">URL Website</label>
                    <Input
                        bind:value={urlwebsite_field}
                        class="required"
                        type="text"
                        placeholder="Url Website"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Link Image</label>
                    <Input
                        bind:value={urlbanner_field}
                        class="required"
                        type="text"
                        placeholder="Link Image"/>
                </div>
            </div>
            <div class="col-lg-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Device</label>
                    <select class="form-control required" bind:value="{devicebanner_field}">
                        <option value="WEB">WEB</option>
                        <option value="MOBILE">MOBILE</option>
                        <option value="DEVICE">DEVICE</option>
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Posisi</label>
                    <select class="form-control required" bind:value="{posisibanner_field}">
                        <option value="TOP">TOP</option>
                        <option value="LEFT">LEFT</option>
                        <option value="RIGHT">RIGHT</option>
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Display</label>
                    <Input
                      bind:value={displaybanner_field}
                      on:keyup={handleKeyboard_format}
                      class="required"
                      type="text"
                      style="text-align:right;"
                      placeholder="Display"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select class="form-control required" bind:value="{statusbanner_field}">
                        <option value="Y">SHOW</option>
                        <option value="N">HIDE</option>
                    </select>
                </div>
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



