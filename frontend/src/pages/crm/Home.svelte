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
	export let listPage = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
    let title_page = "CRM"
    let sData = "";
    let myModal = "";
    
  
    let listcrmsales = []
    let listemployee = []
    let listisbtv = []
    let listPage_isbtv = []
    let totalrecord_isbtv = 0;
    let perpage_isbtv = 0;
    let paging_ibstv = 0;
    let totalpaging_isbtv = 0;
    let pageisbtv_field = 0;
    let pagingnow = 0;
    let buttondownload_isbtv_flag = false;
    let title_modal = "";
    let switchsource_path = "";
    let switchsource_tipe = "";
    let employee_field = "";
    let member_field = "";
   
    let total_sales = 0;
    let field_idrecord = 0;
    let field_nama = "";
    let field_phone = "";
    let field_phone_flag = false;
    let field_status = "";
    let searchcrm = "";
    let filtercrm = "";
    

    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchcrm) {
            filtercrm = listHome.filter(
                (item) =>
                    item.crm_name
                        .toLowerCase()
                        .includes(searchcrm.toLowerCase()) || 
                    item.crm_phone
                        .toLowerCase()
                        .includes(searchcrm.toLowerCase()) || 
                    item.crm_status
                        .toLowerCase()
                        .includes(searchcrm.toLowerCase())
            );
        } else {
            filtercrm = [...listHome];
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const handleSelectPaging = (event) => {
        let page = event.target.value
        pagingnow = page
        const movie = {
                page,
        };
        dispatch("handlePaging", movie);
    };
    const NewData = (e,id,nama,phone,status,totalsales) => {
        sData = e
        if(sData == "New"){
            clearfield_user()
        }else{
            field_phone_flag = true;
            field_idrecord = parseInt(id);
            field_nama = nama;
            field_phone = phone;
            field_status = status;
            total_sales = totalsales
        }
        myModal = new bootstrap.Modal(document.getElementById("modalcruduser"));
        myModal.show();
        
    };
    const DistribusiSales = (e,idcrm,nama,phone,status) => {
        call_employeedepart()
        call_crmsales(phone)
        sData = e
        field_idrecord = parseInt(idcrm);
        field_nama = nama;
        field_phone = phone;
        field_status = status;

        member_field = phone;

        myModal = new bootstrap.Modal(document.getElementById("modalcrmmapping"));
        myModal.show();
        
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(field_nama == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(field_phone == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(field_status == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(field_idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(field_nama == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(field_phone == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(field_status == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/crmsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CRM-SAVE",
                    crm_id: parseInt(field_idrecord),
                    crm_page: pagingnow,
                    crm_phone: field_phone.trim(),
                    crm_name: field_nama,
                    crm_status: field_status,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                if(sData=="New"){
                    clearfield_user()
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
    async function handleSaveStatus(idcrm,status) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/crmsavestatus", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"CRM-SAVE",
                crm_id: parseInt(idcrm),
                crm_page: parseInt(pagingnow),
                crm_status: status,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            
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
    }
    async function handleSave_crmsales() {
        let flag = true
        let msg = ""
        if(member_field == ""){
            flag = false
            msg += "The Member is required\n"
        }
        if(employee_field == ""){
            flag = false
            msg += "The Sales is required\n"
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/crmsalessave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"CRM-SAVE",
                    crm_page: parseInt(pagingnow),
                    search: searchcrm,
                    crmsales_phone: member_field.trim(),
                    crmsales_username: employee_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                if(sData=="New"){
                    clearfield_user()
                }
                msgloader = json.message;
                RefreshHalaman()
                call_crmsales(member_field)
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
    async function handleDelete_crmsales(idcrmsales,phone) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/crmsalesdelete", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page:"CRM-SAVE",
                search: searchcrm,
                crm_page: parseInt(pagingnow),
                crmsales_id: parseInt(idcrmsales),
                crmsales_phone: phone.trim(),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            RefreshHalaman()
            call_crmsales(phone)
        } else if(json.status == 403){
            alert(json.message)
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    const ShowSOURCE = (e) => {
        title_modal = e
        myModal = new bootstrap.Modal(document.getElementById("modalisbtv"));
        myModal.show();
        if(e == "ISBTV"){
            switchsource_path ="/api/crmisbtv"
            switchsource_tipe ="ISBTV"
        }else{
            switchsource_path ="/api/crmduniafilm"
            switchsource_tipe ="DUNIAFILM"
        }
        call_isbtv(switchsource_path,switchsource_tipe)
    };
    async function call_employeedepart() {
        listemployee = []
        const res = await fetch("/api/employeedepart", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page:"CRM-VIEW",
                employee_iddepart:"SLS"
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            let record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                for (var i = 0; i < record.length; i++) {
                    listemployee = [
                        ...listemployee,
                        {
                            employee_username: record[i]["employee_username"],
                            employee_name: record[i]["employee_name"],
                        },
                    ];
                }
            }
        } 
    }
    async function call_crmsales(phone) {
        listcrmsales = []
        const res = await fetch("/api/crmsales", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page:"CRM-VIEW",
                crmsales_phone:phone
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            let record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                for (var i = 0; i < record.length; i++) {
                    listcrmsales = [
                        ...listcrmsales,
                        {
                            crmsales_id: record[i]["crmsales_id"],
                            crmsales_phone: record[i]["crmsales_phone"],
                            crmsales_namamember: record[i]["crmsales_namamember"],
                            crmsales_username: record[i]["crmsales_username"],
                            crmsales_nameemployee: record[i]["crmsales_nameemployee"],
                            crmsales_create: record[i]["crmsales_create"],
                            crmsales_update: record[i]["crmsales_update"],
                        },
                    ];
                }
            }
        } 
    }
    async function call_isbtv(e,type){
        listisbtv = []
        const res = await fetch(e, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                crmisbtv_search: "",
                crmisbtv_page : parseInt(paging_ibstv)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            perpage_isbtv = json.perpage;
            totalrecord_isbtv = json.totalrecord;
            let no = 0;
            if(paging_ibstv > 1){
                no = parseInt(paging_ibstv) 
            }
            if (record != null) {
                totalpaging_isbtv = Math.ceil(parseInt(totalrecord_isbtv) / parseInt(perpage_isbtv))
                if(type=="ISBTV"){
                    for (var i = 0; i < record.length; i++) {
                        let temp_1 = record[i]["crmisbtv_username"];
                        let temp2_1 = temp_1.replace(" ", "");
                        let temp3_1 = temp2_1.replace("-", "");
                        let temp4_1 = temp3_1.replace("(", "");
                        let temp5_1 = temp4_1.replace(")", "");
                        let temp6_1 = temp5_1.replace(" ", "");
                        no = parseInt(no) + 1;
                        listisbtv = [
                            ...listisbtv,
                            {
                                crmisbtv_no: no,
                                crmisbtv_username: temp6_1,
                                crmisbtv_name: record[i]["crmisbtv_name"],
                            },
                        ];
                    }
                }else{
                    for (var i = 0; i < record.length; i++) {
                        let temp = record[i]["crmduniafilm_username"];
                        let temp2 = temp.replace(" ", "");
                        let temp3 = temp2.replace("-", "");
                        let temp4 = temp3.replace("(", "");
                        let temp5 = temp4.replace(")", "");
                        let temp6 = temp5.replace(" ", "");
                        no = parseInt(no) + 1;
                        listisbtv = [
                            ...listisbtv,
                            {
                                crmisbtv_no: no,
                                crmisbtv_username: temp6,
                                crmisbtv_name: record[i]["crmduniafilm_name"],
                            },
                        ];
                    }
                }
                listPage_isbtv = [];
                for(var i=1;i<totalpaging_isbtv;i++){
                    listPage_isbtv = [
                        ...listPage_isbtv,
                        {
                            page_id: i,
                            page_value: ((i*perpage_isbtv)-perpage_isbtv),
                            page_display: i + " Of " + perpage_isbtv*i,
                        },
                    ];
                }
            }
        }
    }
    async function handleDownloadISBTV() {
        console.log(listisbtv.length)
        let flag = true
        let msg = ""
        if(listisbtv.length < 1){
            flag = false
            msg += "The ISBTV is required\n"
        }
        
        if(flag){
            buttondownload_isbtv_flag = true
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/crmsavesource", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "New",
                    page:"CRM-SAVE",
                    crm_page: parseInt(paging_ibstv),
                    crm_source: "ISBTV",
                    crm_data: listisbtv,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            buttondownload_isbtv_flag = false
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "SAVE_USER":
                handleSave();
                break;
            case "SAVE_CRMSALES":
                handleSave_crmsales();
                break;
            case "CALL_ISBTV":
                ShowSOURCE("ISBTV");break;
            case "CALL_DUNIAFILM":
                ShowSOURCE("DUNIAFILM");break;
            case "REFRESH":
                RefreshHalaman();break;
        }
    }
 
    function clearfield_user(){
        field_idrecord = 0;
        field_nama = "";
        field_phone = "";
        field_status = "";
        field_phone_flag = false;
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterIsbtv = [];
                listHome = [];
                const news = {
                    searchIsbtv,
                };
                dispatch("handleNews", news);
        }  
    };
    const handleKeyboard_format = (e) => {
        let numbera;
		for (let i = 0; i < field_phone.length; i++) {
			numbera = parseInt(field_phone[i]);
			if (isNaN(numbera)) {
				if (field_phone[i] != "+") {
					field_phone = "";
				}
			}
		}
        
    };
    const handleSelectGetISBTV = (event) => {
        paging_ibstv = event.target.value
        call_isbtv(switchsource_path,switchsource_tipe)
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
                button_css="btn-primary"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Button
                on:click={callFunction}
                button_function="CALL_ISBTV"
                button_title="Source ISBTV"
                button_css="btn-primary"/>
            <Button
                on:click={callFunction}
                button_function="CALL_DUNIAFILM"
                button_title="Source DUNIA FILM"
                button_css="btn-primary"/>
            &nbsp;&nbsp;&nbsp;
            <Button
                on:click={callFunction}
                button_function="CALL_DUNIAFILM"
                button_title="PROCESS"
                button_css="btn-primary"/>
            <Button
                on:click={callFunction}
                button_function="CALL_DUNIAFILM"
                button_title="INVALID"
                button_css="btn-danger"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-title">
                    <div class="float-end">
                        <select
                            on:change={handleSelectPaging}
                            style="text-align: center;" 
                            class="form-control">
                            {#each listPage as rec}
                                <option value="{rec.page_value}">{rec.page_display}</option>
                            {/each}
                        </select>
                    </div>
                </slot:template>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchcrm}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Username"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                        <table class="table table-striped table-hover table-sm">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">SOURCE</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TEAM SALES</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREATE</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            {#if totalrecord > 0}
                            <tbody>
                                {#each filtercrm as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    NewData("Edit",rec.crm_id,rec.crm_name, rec.crm_phone,rec.crm_status,rec.crm_totalpic);
                                                }} 
                                                class="bi bi-pencil"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    DistribusiSales("Edit",rec.crm_id,rec.crm_name, rec.crm_phone,rec.crm_status);
                                                }} 
                                                class="bi bi-person"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.news_no}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;">
                                            {#if rec.crm_totalpic > 0}
                                                <button
                                                    on:click={() => {
                                                        handleSaveStatus(rec.crm_id,"PROCESS");
                                                    }}  
                                                    type="button" class="btn btn-warning btn-sm">Process</button>
                                            {/if}
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.crm_statuscss}">{rec.crm_status}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                            <a href="https://wa.me/{rec.crm_phone}" target="_blank">{rec.crm_phone}</a>
                                            &nbsp;
                                        </td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_name}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crm_source}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                            {#if rec.crm_pic != null}
                                                {#each rec.crm_pic as rec2 }
                                                    {rec2.crmsales_nameemployee}<br>
                                                {/each}
                                            {/if}
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crm_create}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crm_update}</td>
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
	modal_id="modalcruduser"
	modal_size="modal-dialog-centered"
	modal_title="USER/{sData}"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
			<Input
                bind:value={field_nama}
                class="required"
                type="text"
                placeholder="Name"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Phone</label>
			<Input
                bind:value={field_phone}
                on:keyup={handleKeyboard_format}
                disabled='{field_phone_flag}'
                minlength="6"
				maxlength="20"
                class="required"
                type="text"
                placeholder="Phone"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
			<select class="form-control required" bind:value={field_status}>
                <option value="NEW">NEW</option>
                <option value="INVALID">INVALID</option>
            </select>
		</div>
	</slot:template>
	<slot:template slot="footer">
        {#if total_sales < 1}
        <Button
            on:click={callFunction}
            button_function="SAVE_USER"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>
<Modal
	modal_id="modalisbtv"
	modal_size="modal-dialog-centered"
	modal_title="{title_modal}"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <div style="padding: 10px;">
            <select
                on:change={handleSelectGetISBTV} 
                class="form-control" bind:value="{pageisbtv_field}">
                {#each listPage_isbtv as rec}
                <option value="{rec.page_value}">{rec.page_display}</option>
                {/each}
            </select>
        </div>
	</slot:template>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                </tr>
            </thead>
            <tbody>
                {#each listisbtv as rec }
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crmisbtv_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmisbtv_name}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmisbtv_username}</td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <button
            on:click={() => {
                handleDownloadISBTV();
            }}  
            disabled='{buttondownload_isbtv_flag}' 
            type="button" class="btn btn-warning">DOWNLOAD</button>
	</slot:template>
</Modal>

<Modal
	modal_id="modalcrmmapping"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="Informasi - Member"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input
                        bind:value={field_nama}
                        class="required"
                        disabled
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone</label>
                    <Input
                        bind:value={field_phone}
                        on:keyup={handleKeyboard_format}
                        disabled
                        minlength="6"
                        maxlength="20"
                        class="required"
                        type="text"
                        placeholder="Phone"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select class="form-control required" bind:value={field_status} disabled>
                        <option value="NEW">NEW</option>
                        <option value="VALID">VALID</option>
                        <option value="INVALID">INVALID</option>
                    </select>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="card-header" style="padding: 0px;">
                    <div class="input-group mb-3">
                        <select 
                            class="form-control required" 
                            bind:value={employee_field} >
                            {#each listemployee as rec}
                                <option value="{rec.employee_username}">{rec.employee_name}</option>
                            {/each}
                        </select>
                        <Button
                            on:click={callFunction}
                            button_function="SAVE_CRMSALES"
                            button_title="Save"
                            button_css="btn-warning"/>
                    </div>
                </div>
                <div class="card">
                    <div class="card-body">
                        <table class="table table-striped">
                            <thead>
                                <tr>
                                    <th width="1%" style="text-align: center;vertical-align: top;font-size: 12px;">#</th>
                                    <th width="*" style="text-align: left;vertical-align: top;font-size: 12px;">Sales</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listcrmsales as rec}
                                <tr>
                                    <td>
                                        <i 
                                            style="cursor:pointer;"
                                            on:click={() => {
                                                handleDelete_crmsales(rec.crmsales_id,rec.crmsales_phone);
                                            }} 
                                            class="bi bi-trash"></i>
                                    </td>
                                    <td style="font-size: 12px;">{rec.crmsales_nameemployee}</td>
                                </tr>
                                {/each}
                                
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
	</slot:template>
    <slot:template slot="footer">
        <button
            on:click={() => {
                handleSaveStatus(field_idrecord,"PROCESS");
            }}  
            type="button" class="btn btn-warning ">Process</button>
    </slot:template>
</Modal>