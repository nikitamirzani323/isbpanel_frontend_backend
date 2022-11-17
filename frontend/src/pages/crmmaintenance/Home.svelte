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
    let title_page = "CRM MAINTENANCE"
    let sData = "";
    let myModal = "";
    let files;
  
    
    let listcrmprocess = []
    let listcrmsales = []
    let listcrmdeposit = []
    let listsalesperform = []
    let listsalesperform_noanswer = []
    let listsalesperform_invalid = []
    let listsalesperform_date = []
    let listsalesperform_noanswer_date = []
    let listsalesperform_invalid_date = []
    let listemployee = []
    let listmaintenance = []
    let listdatabase = []
    let listPage_maintenance = []
    let totalrecord_maintenance = 0;
    let totalrecord_sales = 0;
    let perpage_maintenance = 0;
    let paging_maintenance = 0;
    let totalpaging_maintenance = 0;
    let pageimaintenance_field = 0;
    let pagingnow = 0;
    let buttondownload_isbtv_flag = false;
    let title_modal = "";
    let switchsource_path = "";
    let switchsource_tipe = "";
    let employee_field = "";
    let member_field = "";
   
    let title_crmstatus = "";
    let total_crmprocess = 0;
    let total_crmdeposit = 0;
    let total_sales = 0;
    let field_idrecord = 0;
    let field_nama = "";
    let field_phone = "";
    let field_phone_flag = false;
    let field_status = "";
    let searchcrm = "";
    let filtercrm = "";
    let searchcrm_process = "";
    let filtercrm_process = "";

    let css_loader = "display: none;";
    let msgloader = "";

    let info_phone = ""
    let info_nama = ""
    let info_sales = ""
    let info_webagen = ""
    let info_iduseragen = ""
    let info_deposit = 0

    let sales_username = ""
    let sales_nama = ""
    let sales_deposit = 0
    let sales_depositsum = 0
    let sales_noanswer = 0
    let sales_rejected = 0
    let sales_invalid = 0
    let sales_deposit_date = 0
    let sales_depositsum_date = 0
    let sales_noanswer_date = 0
    let sales_rejected_date = 0
    let sales_invalid_date = 0

    let tanggal_start_sales = ""
    let tanggal_end_sales = ""
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
    const infodeposit = (idcrmsales,phone,nama,sales,nmwebagen,idwebagen,deposit) => {
        
        info_phone = phone
        info_nama = nama
        info_sales = sales
        info_webagen = nmwebagen
        info_iduseragen = idwebagen
        info_deposit = deposit
        myModal = new bootstrap.Modal(document.getElementById("modalinfodeposit"));
        myModal.show();
        
    };
    const showProfileSales = (username,name) => {
        sales_nama = name
        sales_username = username
        call_salesperform(username)
        myModal = new bootstrap.Modal(document.getElementById("modalinfosales"));
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
    async function handleUploadDatabase() {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const formData = new FormData();
        formData.append("sdata", sData);
        formData.append("page", "MOVIE-SAVE");
        formData.append("file", files[0]);
        const res = await fetch("/api/crmuploaddatabase", {
            method: "POST",
            headers: {
                Authorization: "Bearer " + token,
            },
            body: formData,
        });
        const json = await res.json();
        let record = json.record;
       
        if (record != null) {
            for (var i = 0; i < record.length; i++) {
                listdatabase = [
                        ...listdatabase,
                        {
                            database_phone: record[i]["Phone"],
                            database_nama: record[i]["Nama"],
                        },
                    ];
            }
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    
    const ShowSOURCE = (e) => {
        title_modal = e
        myModal = new bootstrap.Modal(document.getElementById("modalisbtv"));
        myModal.show();
        
        call_datavalid()
    };
    async function call_salesperform(username) {
        listsalesperform = []
        listsalesperform_noanswer = []
        listsalesperform_invalid = []
        const res = await fetch("/api/crmsalesperform", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page:"CRM-VIEW",
                employee_iddepart:"SLS",
                employee_username:username,
                employee_startdate:"",
                employee_enddate:""
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                totalrecord_sales = record.length;
                
                for (var i = 0; i < record.length; i++) {
                    sales_deposit = record[i]["sales_deposit"]
                    sales_depositsum = record[i]["sales_depositsum"]
                    sales_noanswer = record[i]["sales_noanswer"]
                    sales_rejected = record[i]["sales_reject"]
                    sales_invalid = record[i]["sales_invalid"]
                    if(record[i]["sales_listdeposit"] != null){
                        listsalesperform = record[i]["sales_listdeposit"]
                    }
                    if(record[i]["sales_listnoanswer"] != null){
                        listsalesperform_noanswer = record[i]["sales_listnoanswer"]
                    }
                    if(record[i]["sales_listinvalid"] != null){
                        listsalesperform_invalid = record[i]["sales_listinvalid"]
                    }
                }
            }
        } 
    }
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
                totalrecord_sales = record.length;
                let deposit_css = "";
                let noanswer_css = "";
                let reject_css = "";
                let invalid_css = "";
                
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["employee_deposit"] > 0){
                        deposit_css = "color:blue;font-weight:bold;"
                    }else{
                        deposit_css = "color:red;font-weight:bold;"
                    }
                    if(record[i]["employee_noanswer"] > 0){
                        noanswer_css = "color:blue;font-weight:bold;"
                    }else{
                        noanswer_css = "color:red;font-weight:bold;"
                    }
                    if(record[i]["employee_reject"] > 0){
                        reject_css = "color:blue;font-weight:bold;"
                    }else{
                        reject_css = "color:red;font-weight:bold;"
                    }
                    if(record[i]["employee_invalid"] > 0){
                        invalid_css = "color:blue;font-weight:bold;"
                    }else{
                        invalid_css = "color:red;font-weight:bold;"
                    }
                    listemployee = [
                        ...listemployee,
                        {
                            employee_username: record[i]["employee_username"],
                            employee_name: record[i]["employee_name"],
                            employee_deposit: record[i]["employee_deposit"],
                            employee_depositcss: deposit_css,
                            employee_invalid: record[i]["employee_invalid"],
                            employee_invalidcss: invalid_css,
                            employee_noanswer: record[i]["employee_noanswer"],
                            employee_noanswercss: noanswer_css,
                            employee_reject: record[i]["employee_reject"],
                            employee_rejectcss: reject_css,
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
                crmsales_phone:phone,
                crmsales_status:"MAINTENANCE"
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
    
    async function call_datavalid(){
        listmaintenance = []
        const res = await fetch("/api/crm", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                crm_status: "VALIDMAINTENANCE",
                crm_search: "",
                crm_page : parseInt(paging_maintenance)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            perpage_maintenance = json.perpage;
            totalrecord_maintenance = json.totalrecord;
            let no = 0;
            if(paging_maintenance > 1){
                no = parseInt(paging_maintenance) 
            }
            if (record != null) {
                totalpaging_maintenance = Math.ceil(parseInt(totalrecord_maintenance) / parseInt(perpage_maintenance))
                for (var i = 0; i < record.length; i++) {
                    no = parseInt(no) + 1;
                    listmaintenance = [
                        ...listmaintenance,
                        {
                            crm_no: no,
                            crm_id: record[i]["crm_id"],
                            crm_phone: record[i]["crm_phone"],
                            crm_name: record[i]["crm_name"],
                            crm_pic: record[i]["crm_pic"],
                            crm_totalpic: record[i]["crm_totalpic"],
                            crm_source: record[i]["crm_source"],
                            crm_status: record[i]["crm_status"],
                            crm_statuscss: record[i]["crm_statuscss"],
                            crm_create: record[i]["crm_create"],
                            crm_update: record[i]["crm_update"],
                        },
                    ];
                }
                listPage_maintenance = [];
                for(var i=1;i<totalpaging_maintenance;i++){
                    listPage_maintenance = [
                        ...listPage_maintenance,
                        {
                            page_id: i,
                            page_value: ((i*perpage_maintenance)-perpage_maintenance),
                            page_display: i + " Of " + perpage_maintenance*i,
                        },
                    ];
                }
            }
        }
    }
    async function call_teamsales(){
        call_employeedepart()
        myModal = new bootstrap.Modal(document.getElementById("modallistsales"));
        myModal.show();
    }
    async function handleDownloadMaintenance() {
        let flag = true
        let msg = ""
        if(listmaintenance.length < 1){
            flag = false
            msg += "The MAINTENANCE is required\n"
        }
        
        if(flag){
            buttondownload_isbtv_flag = true
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/crmsavemaintenance", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "New",
                    page:"CRM-SAVE",
                    crm_page: parseInt(paging_maintenance),
                    crm_data: listmaintenance,
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
   
    async function call_crmbystatus(e){
        title_crmstatus = ""
        listcrmprocess = []
        switch(e){
            case "FOLLOWUP":
                title_crmstatus = "CRM - FOLLOWUP"
                break;
            case "VALID":
                title_crmstatus = "CRM - VALID"
                break;
            case "INVALID":
                title_crmstatus = "CRM - INVALID"
                break;
        }
        if(e == "VALID"){
            myModal = new bootstrap.Modal(document.getElementById("modallistcrmvalid"));
            myModal.show();
        }else{
            myModal = new bootstrap.Modal(document.getElementById("modallistcrmstatus"));
            myModal.show();
        }
        
        const res = await fetch("/api/crm", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                crm_status: e,
                crm_search: "",
                crm_page : 0
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            let no = 0;
            
            if (record != null) {
                total_crmprocess = record.length
                for (var i = 0; i < record.length; i++) {
                    no = parseInt(no) + 1;
                    listcrmprocess = [
                        ...listcrmprocess,
                        {
                            crm_no: no,
                            crm_id: record[i]["crm_id"],
                            crm_phone: record[i]["crm_phone"],
                            crm_name: record[i]["crm_name"],
                            crm_pic: record[i]["crm_pic"],
                            crm_totalpic: record[i]["crm_totalpic"],
                            crm_source: record[i]["crm_source"],
                            crm_status: record[i]["crm_status"],
                            crm_statuscss: record[i]["crm_statuscss"],
                            crm_create: record[i]["crm_create"],
                            crm_update: record[i]["crm_update"],
                        },
                    ];
                }
            }
        }
    }
    
    function callFunction(event){
        switch(event.detail){
            case "CALL_CRMFOLLOWUP":
                call_crmbystatus("FOLLOWUP");
                break;
            case "CALL_CRMVALID":
                call_crmbystatus("VALID");
                break;
            case "CALL_CRMINVALID":
                call_crmbystatus("INVALID");
                break;
            case "CALL_VALID":
                ShowSOURCE();
                break;
            case "SAVE_USER":
                handleSave();
                break;
            case "SAVE_CRMSALES":
                handleSave_crmsales();
                break;
            case "UPLOAD_DATABASE":
                handleUploadDatabase();
                break;
            case "CALL_TEAMSALES":
                call_teamsales();break;
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
    const handleSelectGetMaintenance = (event) => {
        paging_maintenance = event.target.value
        call_datavalid()
    };
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

        if (searchcrm_process) {
            filtercrm_process = listcrmprocess.filter(
                (item) =>
                    item.crm_name
                        .toLowerCase()
                        .includes(searchcrm_process.toLowerCase()) || 
                    item.crm_phone
                        .toLowerCase()
                        .includes(searchcrm_process.toLowerCase()) 
            );
        } else {
            filtercrm_process = [...listcrmprocess];
        }
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
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Button
                on:click={callFunction}
                button_function="CALL_VALID"
                button_title="Source VALID"
                button_css="btn-primary"/>
            &nbsp;&nbsp;&nbsp;
            <Button
                on:click={callFunction}
                button_function="CALL_TEAMSALES"
                button_title="TEAM SALES"
                button_css="btn-primary"/>
            &nbsp;&nbsp;&nbsp;
            <Button
                on:click={callFunction}
                button_function="CALL_CRMFOLLOWUP"
                button_title="FOLLOWUP"
                button_css="btn-warning"/>
            <Button
                on:click={callFunction}
                button_function="CALL_CRMVALID"
                button_title="VALID"
                button_css="btn-success"/>
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
                            placeholder="Search Phone"
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
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crm_no}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;">
                                            {#if rec.crm_totalpic > 0}
                                                <button
                                                    on:click={() => {
                                                        handleSaveStatus(rec.crm_id,"FOLLOWUP");
                                                    }}  
                                                    type="button" class="btn btn-warning btn-sm">Follow Up</button>
                                            {/if}
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.crm_statuscss}">
                                                {rec.crm_status}
                                            </span>
                                        </td>
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
	modal_title="INFORMASI MEMBER"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
			<Input
                bind:value={field_nama}
                disabled 
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
	</slot:template>
	<slot:template slot="footer">
	</slot:template>
</Modal>
<Modal
	modal_id="modalisbtv"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="SOURCE VALID"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <div style="padding: 10px;">
            <select
                on:change={handleSelectGetMaintenance} 
                class="form-control" bind:value="{pageimaintenance_field}">
                {#each listPage_maintenance as rec}
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
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th width="20%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TEAM SALES</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                </tr>
            </thead>
            <tbody>
                {#each listmaintenance as rec }
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crm_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_name}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_phone}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                        {#each rec.crm_pic as rec2}
                            {rec2.crmsales_nameemployee}  - 
                            {#if rec2.crmsales_status == "DEPOSIT"}
                                <span style="padding:5px;padding-left: 10px;padding-right: 10px;background:#ffc107;border-radius: 50px;">{rec2.crmsales_status}</span>
                            {/if}
                            {#if rec2.crmsales_status == "REJECT" || rec2.crmsales_status=="NOANSWER"}
                                <span style="padding:5px;padding-left: 10px;padding-right: 10px;background:#dc3545;border-radius: 50px;color:white;">{rec2.crmsales_status}</span>
                            {/if}
                            - 
                            {#if rec2.crmsales_status == "DEPOSIT"}
                                <i 
                                    on:click={() => {
                                        infodeposit(rec2.crmsales_idcrmsales,rec.crm_phone,rec.crm_name,rec2.crmsales_nameemployee,rec2.crmsales_nmwebagen,rec2.crmsales_idwebagen,rec2.crmsales_deposit);
                                    }} 
                                    class="bi bi-info-circle"  style="cursor:pointer;"></i>
                            {/if}
                            {#if rec2.crmsales_note != ""}
                                <i class="bi bi-chat-left-dots" title="{rec2.crmsales_note}" style="cursor:pointer;"></i>
                            {/if}
                            <br>
                        {/each} 
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_update}</td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <button
            on:click={() => {
                handleDownloadMaintenance();
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
                handleSaveStatus(field_idrecord,"FOLLOWUP");
            }}  
            type="button" class="btn btn-warning ">Follow Up</button>
    </slot:template>
</Modal>

<Modal
	modal_id="modallistcrmstatus"
	modal_size="modal-dialog-centered modal-xl"
	modal_title="{title_crmstatus}"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <div class="col-lg-12" style="padding: 5px;">
            <input
                bind:value={searchcrm_process}
                type="text"
                class="form-control"
                placeholder="Search Phone and Name"
                aria-label="Search"/>
        </div>
	</slot:template>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TEAM SALES</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CREATE</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                </tr>
            </thead>
            <tbody>
                {#each filtercrm_process as rec }
                <tr>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_phone}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_name}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                        {#each rec.crm_pic as rec2}
                            {#if rec2.crmsales_status_utama == ""}
                            {rec2.crmsales_nameemployee}<br>
                            {/if}
                        {/each} 
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_create}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_update}</td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <b>TOTAL : {total_crmprocess}</b>
	</slot:template>
</Modal>
<Modal
	modal_id="modallistcrmvalid"
	modal_size="modal-dialog-centered modal-xl"
	modal_title="{title_crmstatus}"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <div class="col-lg-12" style="padding: 5px;">
            <input
                bind:value={searchcrm_process}
                type="text"
                class="form-control"
                placeholder="Search Phone and Name"
                aria-label="Search"/>
        </div>
	</slot:template>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th width="20%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TEAM SALES</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CREATE</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                </tr>
            </thead>
            <tbody>
                {#each filtercrm_process as rec }
                <tr>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_phone}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_name}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                        {#each rec.crm_pic as rec2}
                            {rec2.crmsales_nameemployee}  - 
                            {#if rec2.crmsales_status == "DEPOSIT"}
                                <span style="padding:5px;padding-left: 10px;padding-right: 10px;background:#ffc107;border-radius: 50px;">{rec2.crmsales_status}</span>
                            {/if}
                            {#if rec2.crmsales_status == "REJECT" || rec2.crmsales_status=="NOANSWER"}
                                <span style="padding:5px;padding-left: 10px;padding-right: 10px;background:#dc3545;border-radius: 50px;color:white;">{rec2.crmsales_status}</span>
                            {/if}
                            - 
                            {#if rec2.crmsales_status == "DEPOSIT"}
                                <i 
                                    on:click={() => {
                                        infodeposit(rec2.crmsales_idcrmsales,rec.crm_phone,rec.crm_name,rec2.crmsales_nameemployee,rec2.crmsales_nmwebagen,rec2.crmsales_idwebagen,rec2.crmsales_deposit);
                                    }} 
                                    class="bi bi-info-circle"  style="cursor:pointer;"></i>
                            {/if}
                            {#if rec2.crmsales_note != ""}
                                <i class="bi bi-chat-left-dots" title="{rec2.crmsales_note}" style="cursor:pointer;"></i>
                            {/if}
                            <br>
                        {/each} 
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_create}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crm_update}</td>
                </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <b>TOTAL : {total_crmprocess}</b>
	</slot:template>
</Modal>

<Modal
	modal_id="modalinfodeposit"
	modal_size="modal-dialog-centered "
	modal_title="DEPOSIT"
    modal_body_css="height:400px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-sm">
            <tbody>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">PHONE</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">{info_phone}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">NAME</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">{info_nama}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">TEAM SALES</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">{info_sales}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">WEBSITE</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">{info_webagen}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">IDUSER</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">{info_iduseragen}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">DEPOSIT</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">
                        {new Intl.NumberFormat().format(info_deposit)}
                    </td>
                </tr>
            </tbody>
        </table>
	</slot:template>
</Modal>

<Modal
	modal_id="modallistsales"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="TEAM SALES"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={false}
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th width="15%" style="text-align:right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DEPOSIT</th>
                    <th width="15%" style="text-align:right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NOANSWER</th>
                    <th width="15%" style="text-align:right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">REJECT</th>
                    <th width="15%" style="text-align:right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">INVALID</th>
                </tr>
            </thead>
            <tbody>
                {#each listemployee as rec}
                    <tr>
                        <td on:click={() => {
                            showProfileSales(rec.employee_username,rec.employee_name);
                        }} NOWRAP style="text-align:left;vertical-align: top;font-size: {table_body_font};text-decoration: underline;cursor:pointer;">
                            {rec.employee_name}
                        </td>
                        <td NOWRAP style="text-align:right;vertical-align: top;font-size: {table_body_font};{rec.employee_depositcss}">{rec.employee_deposit}</td>
                        <td NOWRAP style="text-align:right;vertical-align: top;font-size: {table_body_font};{rec.employee_noanswercss}">{rec.employee_noanswer}</td>
                        <td NOWRAP style="text-align:right;vertical-align: top;font-size: {table_body_font};{rec.employee_rejectcss}">{rec.employee_reject}</td>
                        <td NOWRAP style="text-align:right;vertical-align: top;font-size: {table_body_font};{rec.employee_invalidcss}">{rec.employee_invalid}</td>
                    </tr>
                {/each}
                
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        TOTAL  : <span style="color:blue;font-weight:bold;">{new Intl.NumberFormat().format(totalrecord_sales)}</span>
	</slot:template>
</Modal>

<Modal
	modal_id="modalinfosales"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="SALES : {sales_nama}"
    modal_body_css="height:600px;overflow-y: scroll;"
    modal_footer_css="padding:2px;"
	modal_footer={false}>
	<slot:template slot="body">
        <ul class="nav nav-tabs">
            <li class="nav-item">
              <a class="nav-link active" data-bs-toggle="tab" href="#sales_info">INFO</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" data-bs-toggle="tab" href="#sales_bydate">BY DATE</a>
            </li>
        </ul>
        <div id="myTabContentpusat" class="tab-content">
            <div class="tab-pane fade show active" id="sales_info">
                <table class="table table-sm">
                    <tbody>
                        <tr>
                            <td width="20%" style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL DEPOSIT</td>
                            <td width="1%" style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td width="*" style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_deposit)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL NOANSWER</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_noanswer)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL REJECTED</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_rejected)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL INVALID</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_invalid)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">SUBTOTAL DEPOSIT</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">
                                {new Intl.NumberFormat().format(sales_depositsum)}
                            </td>
                        </tr>
                    </tbody>
                </table>
                <ul class="nav nav-tabs">
                    <li class="nav-item">
                      <a class="nav-link active" data-bs-toggle="tab" href="#sales_deposit">DEPOSIT</a>
                    </li>
                    <li class="nav-item">
                      <a class="nav-link" data-bs-toggle="tab" href="#sales_noanswer">NOANSWER / REJECT</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" data-bs-toggle="tab" href="#sales_invalid">INVALID</a>
                    </li>
                </ul>
                <div id="myTabContent" class="tab-content">
                    <div class="tab-pane fade show active" id="sales_deposit">
                        <table class="table table-sm">
                            <thead>
                                <tr>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                                    <th nowrap width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SOURCE</th>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WEBSITE AGEN</th>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ID AGEN</th>
                                    <th nowrap width="10%" style="text-align:right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DEPOSIT</th>
                                    <th nowrap width="20%" style="text-align:center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listsalesperform as rec}
                                    <tr>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_phone}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_nama}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_source}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_nmwebagen}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_iduseragen}</td>
                                        <td nowrap style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(rec.crmdeposit_deposit)}</td>
                                        <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_update}</td>
                                    </tr>
                                {/each}
                                
                            </tbody>
                        </table>
                    </div>
                    <div class="tab-pane fade" id="sales_noanswer">
                        <table class="table table-sm">
                            <thead>
                                <tr>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                                    <th width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SOURCE</th>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TIPE</th>
                                    <th width="20%" style="text-align:center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listsalesperform_noanswer as rec}
                                    <tr>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_phone}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_nama}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_source}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_tipe}</td>
                                        <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_update}</td>
                                    </tr>
                                {/each}
                                
                            </tbody>
                        </table>
                    </div>
                    <div class="tab-pane fade" id="sales_invalid">
                        <table class="table table-sm">
                            <thead>
                                <tr>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                                    <th width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SOURCE</th>
                                    <th width="20%" style="text-align:center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listsalesperform_invalid as rec}
                                    <tr>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_phone}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_nama}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_source}</td>
                                        <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_update}</td>
                                    </tr>
                                {/each}
                                
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
            <div class="tab-pane " id="sales_bydate">
                <div class="input-group mb-4">
                    <Input
                        bind:value={tanggal_start_sales}
                        class="required"
                        type="date"
                        name="date"
                        id="tanggal_start_sales"
                        data-date-format="dd-mm-yyyy"
                        placeholder="date placeholder"/>
                    <span class="input-group-text">s/d</span>
                    <Input
                        bind:value={tanggal_end_sales}
                        class="required"
                        type="date"
                        name="date"
                        id="tanggal_end_sales"
                        data-date-format="dd-mm-yyyy"
                        placeholder="date placeholder"/>
                    <Button
                        on:click={callFunction}
                        button_function="SALESDATEPERFORM"
                        button_title="Generate"
                        button_css="btn-warning"/>
                </div>
                <table class="table table-sm">
                    <tbody>
                        <tr>
                            <td width="20%" style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL DEPOSIT</td>
                            <td width="1%" style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td width="*" style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_deposit_date)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL NOANSWER</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_noanswer_date)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL REJECTED</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_rejected_date)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">TOTAL INVALID</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(sales_invalid_date)}</td>
                        </tr>
                        <tr>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">SUBTOTAL DEPOSIT</td>
                            <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                            <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">
                                {new Intl.NumberFormat().format(sales_depositsum_date)}
                            </td>
                        </tr>
                    </tbody>
                </table>
                <ul class="nav nav-tabs">
                    <li class="nav-item">
                      <a class="nav-link active" data-bs-toggle="tab" href="#sales_deposit_date">DEPOSIT</a>
                    </li>
                    <li class="nav-item">
                      <a class="nav-link" data-bs-toggle="tab" href="#sales_noanswer_date">NOANSWER / REJECT</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" data-bs-toggle="tab" href="#sales_invalid_date">INVALID</a>
                    </li>
                </ul>
                <div id="myTabContent" class="tab-content">
                    <div class="tab-pane fade show active" id="sales_deposit_date">
                        <table class="table table-sm">
                            <thead>
                                <tr>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                                    <th nowrap width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SOURCE</th>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WEBSITE AGEN</th>
                                    <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ID AGEN</th>
                                    <th nowrap width="10%" style="text-align:right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DEPOSIT</th>
                                    <th nowrap width="20%" style="text-align:center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listsalesperform_date as rec}
                                    <tr>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_phone}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_nama}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_source}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_nmwebagen}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_iduseragen}</td>
                                        <td nowrap style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(rec.crmdeposit_deposit)}</td>
                                        <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crmdeposit_update}</td>
                                    </tr>
                                {/each}
                                
                            </tbody>
                        </table>
                    </div>
                    <div class="tab-pane fade" id="sales_noanswer_date">
                        <table class="table table-sm">
                            <thead>
                                <tr>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                                    <th width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SOURCE</th>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TIPE</th>
                                    <th width="20%" style="text-align:center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listsalesperform_noanswer_date as rec}
                                    <tr>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_phone}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_nama}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_source}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_tipe}</td>
                                        <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crmnoanswer_update}</td>
                                    </tr>
                                {/each}
                                
                            </tbody>
                        </table>
                    </div>
                    <div class="tab-pane fade" id="sales_invalid_date">
                        <table class="table table-sm">
                            <thead>
                                <tr>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                                    <th width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                                    <th width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">SOURCE</th>
                                    <th width="20%" style="text-align:center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listsalesperform_invalid_date as rec}
                                    <tr>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_phone}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_nama}</td>
                                        <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_source}</td>
                                        <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crminvalid_update}</td>
                                    </tr>
                                {/each}
                                
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>   
	</slot:template>
</Modal>

<Modal
	modal_id="modaldatabase"
	modal_size="modal-dialog-centered"
	modal_title="Source Database"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <div class="input-group" style="padding: 10px;">
            <input id="fileUpload" type="file" bind:files />
            <Button
                on:click={callFunction}
                button_function="UPLOAD_DATABASE"
                button_title="Upload"
                button_css="btn-warning"/>
        </div>
	</slot:template>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="20%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PHONE</th>
                    <th width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                </tr>
            </thead>
            <tbody>
                {#each listdatabase as rec}
                <tr>
                    <td>{rec.database_phone}</td>
                    <td>{rec.database_nama}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_DATABASE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>