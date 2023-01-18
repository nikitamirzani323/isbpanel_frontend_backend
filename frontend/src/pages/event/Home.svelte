<script>
    import { Input } from "sveltestrap";
    import dayjs from "dayjs";
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
	let title_page = "EVENT"
    let sData = "";
    let sDataNewPartisipasi = "";
    let myModal_newentry = "";
    let myModal_memberagen = "";
    let idevent_global = 0;
    let idwebsite_global = 0;
    let nmwebsite_global = "";
    let idwebsite_field = 0;
    let nmevent_field = "";
    let startevent_field = "";
    let startevent_jam_field = "00:00";
    let endevent_field = "";
    let endevent_jam_field = "00:00";
    let mindeposit_field = 0;
    let create_field = "";
    let update_field = "";

    let listpartisipasi_total = 0;
    let listpartisipasi_db = [];
    let listwebsiteagen_db = [];
    let listmemberagen_db = [];
    
    let idmemberagen_partisipasi_field = 0;
    let username_partisipasi_field = "";
    let deposito_partisipasi_field = 0;

    let searchMember = "";
    let filterMember = [];
    let css_loader = "display: none;";
    let msgloader = "";
    
   
    $: {
        if (searchMember) {
            filterMember = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchMember.toLowerCase())
            );
        } else {
            filterMember = [...listHome];
        }
        
    }
    const NewData = (e,idwebsite,event,start,end,create,update) => {
        sData = e
        call_websiteagen()
        if(sData == "New"){
            clearField()
        }else{
            // alert(idwebsite)
            idwebsite_field = idwebsite;
            nmevent_field = event;
            startevent_field = dayjs(start).format("YYYY-MM-DD");
            startevent_jam_field = dayjs(start).format("HH:MM");
            endevent_field = dayjs(end).format("YYYY-MM-DD");
            endevent_jam_field = dayjs(end).format("HH:MM");
            create_field = create;
            update_field = update;
            
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const ListPartisipasi = (idevent,idwebsite,nmwebsite) => {
        call_listpartisipasi(idevent)
        idevent_global = idevent
        idwebsite_global = idwebsite
        nmwebsite_global = nmwebsite
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistpartisipasi"));
        myModal_newentry.show();
    };
    const NewFormPartisipasi = () => {
        sDataNewPartisipasi = "New";
        clearFieldPartisipasi();
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrypartisipasi"));
        myModal_newentry.show();
    };
    const ListMemberAgen = () => {
        call_memberagen(idwebsite_global)
        myModal_memberagen = new bootstrap.Modal(document.getElementById("modallistmemberagen"));
        myModal_memberagen.show();
    };
    const InsertPartisipasi = (id,e) => {
        idmemberagen_partisipasi_field = parseInt(id);
        username_partisipasi_field = e;
        myModal_memberagen.hide();
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function call_listpartisipasi(idevent) {
        listpartisipasi_db = [];
        listpartisipasi_total = 0;
        const res = await fetch("/api/eventdetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "MOVIEALBUM-VIEW",
                event_id: idevent,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listpartisipasi_total = listpartisipasi_total + record[i]["eventdetail_deposit"];
                    listpartisipasi_db = [
                        ...listpartisipasi_db,
                        {
                            eventdetail_no: no,
                            eventdetail_id: record[i]["eventdetail_id"],
                            eventdetail_phone: record[i]["eventdetail_phone"],
                            eventdetail_username: record[i]["eventdetail_username"],
                            eventdetail_voucher: record[i]["eventdetail_voucher"],
                            eventdetail_deposit: record[i]["eventdetail_deposit"],
                            eventdetail_create: record[i]["eventdetail_create"],
                            eventdetail_update: record[i]["eventdetail_update"],
                        },
                    ];
                }
            }
        }
    }
    async function call_websiteagen() {
        listwebsiteagen_db = [];
        const res = await fetch("/api/webagen", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "MOVIEALBUM-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listwebsiteagen_db = [
                        ...listwebsiteagen_db,
                        {
                            websiteagen_id: record[i]["websiteagen_id"],
                            websiteagen_name: record[i]["websiteagen_name"],
                        },
                    ];
                }
            }
        }
    }
    async function call_memberagen(e) {
        listmemberagen_db = [];
        const res = await fetch("/api/memberselect", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "MOVIEALBUM-VIEW",
                memberagen_idwebagen: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listmemberagen_db = [
                        ...listmemberagen_db,
                        {
                            memberagen_no: no,
                            memberagen_id: record[i]["memberagen_id"],
                            memberagen_username: record[i]["memberagen_username"],
                            memberagen_phone: record[i]["memberagen_phone"],
                            memberagen_name: record[i]["memberagen_name"],
                        },
                    ];
                }
            }
        }
    }
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idwebsite_field == 0){
                flag = false
                msg += "The Website Event is required\n"
            }
            if(nmevent_field == ""){
                flag = false
                msg += "The Name Event is required\n"
            }
            if(startevent_field == ""){
                flag = false
                msg += "The Start Event is required\n"
            }
            if(endevent_field == ""){
                flag = false
                msg += "The End Event is required\n"
            }
        }else{
            if(idwebsite_field == 0){
                flag = false
                msg += "The Website Event is required\n"
            }
            if(nmevent_field == ""){
                flag = false
                msg += "The Name Event is required\n"
            }
            if(startevent_field == ""){
                flag = false
                msg += "The Start Event is required\n"
            }
            if(endevent_field == ""){
                flag = false
                msg += "The End Event is required\n"
            }
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/eventsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"DOMAIN-SAVE",
                    event_idwebagen: idwebsite_field,
                    event_name: nmevent_field,
                    event_startevent: startevent_field+" "+startevent_jam_field,
                    event_endevent: endevent_field+" "+endevent_jam_field,
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
    async function handleSavePartisipasi() {
        let flag = true
        let msg = ""
        if(sDataNewPartisipasi == "New"){
            if(idmemberagen_partisipasi_field == 0){
                flag = false
                msg += "The Member is required\n"
            }
            if(deposito_partisipasi_field == ""){
                flag = false
                msg += "The Deposit is required\n"
            }
        }else{
            if(idmemberagen_partisipasi_field == 0){
                flag = false
                msg += "The Member is required\n"
            }
            if(deposito_partisipasi_field == ""){
                flag = false
                msg += "The Deposit is required\n"
            }
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/eventdetailsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataNewPartisipasi,
                    page:"DOMAIN-SAVE",
                    eventdetail_id: idwebsite_field,
                    eventdetail_idevent: parseInt(idevent_global),
                    eventdetail_idmemberagen: parseInt(idmemberagen_partisipasi_field),
                    eventdetail_deposit: parseInt(deposito_partisipasi_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                
                msgloader = json.message;
                call_listpartisipasi(idevent_global)
                clearFieldPartisipasi();
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
        idevent_global = 0;
        idwebsite_global = 0;
        nmwebsite_global = "";
        idwebsite_field = 0;
        nmevent_field = "";
        startevent_field = "";
        startevent_jam_field = "00:00";
        endevent_field = "";
        endevent_jam_field = "00:00";
        create_field = "";
        update_field = "";
    }
    function clearFieldPartisipasi(){
        idmemberagen_partisipasi_field = 0;
        username_partisipasi_field = "";
        deposito_partisipasi_field = 0;
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
            case "FORM_PARTISIPASI":
                NewFormPartisipasi();break;
            case "SAVENEWPARTISIPASI":
                handleSavePartisipasi();break;
        }
    }
    const handleKeyboard_format = () => {
        let numbera;

        for (let i = 0; i < startevent_jam_field.length; i++) {
            numbera = parseInt(startevent_jam_field[i]);
            if (isNaN(numbera)) {
                if (startevent_jam_field[i] != ":") {
                    startevent_jam_field = "";
                }
            }
        }
        for (let i = 0; i < endevent_jam_field.length; i++) {
            numbera = parseInt(endevent_jam_field[i]);
            if (isNaN(numbera)) {
                if (endevent_jam_field[i] != ":") {
                    endevent_jam_field = "";
                }
            }
        }
        
        for (let i = 0; i < mindeposit_field.length; i++) {
            numbera = parseInt(mindeposit_field[i]);
            if (isNaN(numbera)) {
                mindeposit_field = 0;
            }
        }
        for (let i = 0; i < deposito_partisipasi_field.length; i++) {
            numbera = parseInt(deposito_partisipasi_field[i]);
            if (isNaN(numbera)) {
                deposito_partisipasi_field = 0;
            }
        }
    };
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
                            bind:value={searchMember}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Event"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">WEBSITE AGEN</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAMA</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">START</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">END</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREATE</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">UPDATE</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterMember as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewData("Edit",rec.home_idwebsite, rec.home_name,rec.home_start, rec.home_end,rec.home_create,rec.home_update);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                ListPartisipasi(rec.home_id,rec.home_idwebsite,rec.home_websiteagen);
                                            }} 
                                            class="bi bi-person-badge"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_websiteagen}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_start}</td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_end}</td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_create}</td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_update}</td>
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
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_body_css="height:490px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Website</label>
            <select bind:value={idwebsite_field} class="form-control required">
                {#each listwebsiteagen_db as rec}
                    <option value="{rec.websiteagen_id}">{rec.websiteagen_name}</option>
                {/each}
            </select>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Event</label>
            <Input
                bind:value={nmevent_field}
                class="required"
                type="text"
                placeholder="Event"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Start Event</label>
            <div class="input-group">
                <Input
                    bind:value={startevent_field}
                    class="required"
                    type="date"
                    name="date"
                    id="exampleDate"
                    data-date-format="dd-mm-yyyy"
                    placeholder="date placeholder"/>
                <Input
                    bind:value={startevent_jam_field}
                    on:keyup={handleKeyboard_format}
                    class="required"
                    style="text-align:center;"
                    type="text"
                    placeholder="00:00"/>
            </div>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">End Event</label>
            <div class="input-group">
                <Input
                    bind:value={endevent_field}
                    class="required"
                    type="date"
                    name="date"
                    id="exampleDate"
                    data-date-format="dd-mm-yyyy"
                    placeholder="date placeholder"/>
                <Input
                    bind:value={endevent_jam_field}
                    on:keyup={handleKeyboard_format}
                    class="required"
                    style="text-align:center;"
                    type="text"
                    placeholder="00:00"/>
            </div>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Min Deposit</label>
            <input
                bind:value={mindeposit_field}
                on:keyup={handleKeyboard_format}
                class="form-control required"
                type="text"
                style="text-align:right;"
                placeholder=""/>
            <div id="emailHelp" class="form-text" style="color:blue;text-align: right;font-size: 11px;">
                {new Intl.NumberFormat().format(mindeposit_field)}
            </div>

        </div>
        {#if sData == "Edit"}
            <div class="alert alert-dark" role="alert" style="font-size:11px;padding:5px;">
                Create : {create_field}
                <br />
                Update : {update_field}
            </div>
        {/if}
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
	modal_id="modallistpartisipasi"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="List Partisipasi - {nmwebsite_global}"
    modal_body_css="height:400px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <ul class="nav nav-tabs" role="tablist">
            <li class="nav-item" role="presentation">
              <a class="nav-link active" data-bs-toggle="tab" href="#homepanel" aria-selected="true" role="tab">ListAll</a>
            </li>
            <li class="nav-item" role="presentation">
              <a class="nav-link" data-bs-toggle="tab" href="#grouppanel" aria-selected="false" tabindex="-1" role="tab">Group</a>
            </li>
          </ul>
          <div id="myTabContent" class="tab-content">
            <div class="tab-pane fade show active" id="homepanel" role="tabpanel">
                <table class="table table-striped ">
                    <thead>
                        <tr>
                            <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                            <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                            <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                            <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">VOUCHER</th>
                            <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPOSIT</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each listpartisipasi_db as rec}
                        <tr>
                            <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_no}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_username}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_phone}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_voucher}</td>
                            <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.eventdetail_deposit)}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="tab-pane fade" id="grouppanel" role="tabpanel">
              <p>Food truck fixie locavore, accusamus mcsweeney's marfa nulla single-origin coffee squid. Exercitation +1 labore velit, blog sartorial PBR leggings next level wes anderson artisan four loko farm-to-table craft beer twee. Qui photo booth letterpress, commodo enim craft beer mlkshk aliquip jean shorts ullamco ad vinyl cillum PBR. Homo nostrud organic, assumenda labore aesthetic magna delectus mollit.</p>
            </div>
          </div>
        
	</slot:template>
	<slot:template slot="footer">
        <div style="font-size: 12px;">
            SUBTOTAL : <span style="color:blue;font-weight:bold;">{new Intl.NumberFormat().format(listpartisipasi_total)}</span>
        </div>
        <Button
            on:click={callFunction}
            button_function="FORM_PARTISIPASI"
            button_title="New Partisipasi"
            button_css="btn-warning"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalentrypartisipasi"
	modal_size="modal-dialog-centered"
	modal_title="New Partisipasi"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Member</label>
            <div class="input-group">
                <input 
                    bind:value={idmemberagen_partisipasi_field}
                    type="hidden" >
                <input
                    disabled
                    bind:value={username_partisipasi_field} 
                    type="text" class="form-control required" placeholder="Username">
                <button
                    on:click={() => {
                        ListMemberAgen();
                    }}  
                    class="btn btn-info" type="button">
                    <i class="bi bi-search"></i>
                </button>
                
            </div>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Deposit</label>
            <input
                bind:value={deposito_partisipasi_field}
                on:keyup={handleKeyboard_format}
                class="form-control required"
                type="text"
                style="text-align:right;"
                placeholder=""/>
            <div id="emailHelp" class="form-text" style="color:blue;text-align: right;font-size: 11px;">
                {new Intl.NumberFormat().format(deposito_partisipasi_field)}
            </div>

        </div>
        {#if sData == "Edit"}
            <div class="alert alert-dark" role="alert" style="font-size:11px;padding:5px;">
                Create : {create_field}
                <br />
                Update : {update_field}
            </div>
        {/if}
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVENEWPARTISIPASI"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modallistmemberagen"
	modal_size="modal-dialog-centered"
	modal_title="List Member Agen"
    modal_body_css="height:400px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                </tr>
            </thead>
            <tbody>
                {#each listmemberagen_db as rec}
                <tr 
                    on:click={() => {
                        InsertPartisipasi(rec.memberagen_id,rec.memberagen_username+" - "+rec.memberagen_phone+" - "+rec.memberagen_name);
                    }}  
                    style="text-decoration: underline;color:biru;cursor: pointer;">
                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_no}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_username}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_phone}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_name}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">

        <Button
            on:click={callFunction}
            button_function="FORM_PARTISIPASI"
            button_title="New Partisipasi"
            button_css="btn-warning"/>
	</slot:template>
</Modal>