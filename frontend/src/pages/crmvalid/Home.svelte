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
    let title_page = "CRM VALID"
    let sData = "";
    let myModal = "";
    let files;
  
    
    let listcrmprocess = []
    let listcrmsales_deposit = []
    let listcrmsales_noanswer = []
    let pagingnow = 0;
    let field_nama = "";
    let field_phone = "";
    let field_source = "";
    let searchcrm = "";
    let filtercrm = "";
    let searchcrm_process = "";
    let filtercrm_process = "";

    let css_loader = "display: none;";
    let msgloader = "";

    
    let total_deposit = 0
    let total_depositsum = 0
    let total_noanswer = 0
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
    const NewData = (nama,phone,source,sales) => {
        field_nama = nama;
        field_phone = phone;
        field_source = source;
        total_deposit = 0
        total_depositsum = 0
        total_noanswer = 0
        listcrmsales_deposit = [];
        listcrmsales_noanswer = [];
        for(let i=0;i<sales.length;i++){
            if(sales[i]["crmsales_status"] == "DEPOSIT"){
                listcrmsales_deposit = [
                    ...listcrmsales_deposit,
                    {
                        crmsales_nameemployee: sales[i]["crmsales_nameemployee"],
                        crmsales_nmwebagen: sales[i]["crmsales_nmwebagen"],
                        crmsales_idwebagen: sales[i]["crmsales_idwebagen"],
                        crmsales_deposit: sales[i]["crmsales_deposit"],
                    },
                ];
                total_depositsum = total_depositsum + sales[i]["crmsales_deposit"]
                total_deposit = total_deposit + 1
            }else{
                listcrmsales_noanswer = [
                    ...listcrmsales_noanswer,
                    {
                        crmsales_nameemployee: sales[i]["crmsales_nameemployee"],
                        crmsales_note: sales[i]["crmsales_note"],
                        crmsales_status: sales[i]["crmsales_status"],
                    },
                ];
                total_noanswer = total_noanswer + 1
            }
        }
        myModal = new bootstrap.Modal(document.getElementById("modalcruduser"));
        myModal.show();
        
    };
 
    
    function callFunction(event){
        switch(event.detail){
            case "REFRESH":
                RefreshHalaman();break;
        }
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
                        <table class="table table-striped table-sm">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
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
                                                    NewData(rec.crm_name, rec.crm_phone,rec.crm_source,rec.crm_pic);
                                                }} 
                                                class="bi bi-info-circle"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.crm_no}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;">
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
	modal_size="modal-dialog-centered modal-lg"
	modal_title="INFORMASI MEMBER"
    modal_body_css="height:600px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-sm">
            <tbody>
                <tr>
                    <td width="20%" style="text-align: left;vertical-align: top;font-size: {table_body_font};">Name</td>
                    <td width="1%" style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td width="*" style="text-align: left;vertical-align: top;font-size: {table_body_font};">{field_nama}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">PHONE</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">{field_phone}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">SOURCE</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">{field_source}</td>
                </tr>
                <tr>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">SUBTOTAL DEPOSIT</td>
                    <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">:</td>
                    <td style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">
                        {new Intl.NumberFormat().format(total_depositsum)}
                    </td>
                </tr>
            </tbody>
        </table>
        <ul class="nav nav-tabs">
            <li class="nav-item">
                <a class="nav-link active" data-bs-toggle="tab" href="#sales_deposit">DEPOSIT 
                    ({total_deposit})
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link" data-bs-toggle="tab" href="#sales_noanswer">NOANSWER / REJECT
                    ({total_noanswer})
                </a>
            </li>
        </ul>
        <div id="myTabContent" class="tab-content">
            <div class="tab-pane fade show active" id="sales_deposit">
                <table class="table table-sm">
                    <thead>
                        <tr>
                            <th nowrap width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TEAM SALES</th>
                            <th nowrap width="20%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WEBSITE AGEN</th>
                            <th nowrap width="10%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ID AGEN</th>
                            <th nowrap width="15%" style="text-align:right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DEPOSIT</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each listcrmsales_deposit as rec}
                            <tr>
                                <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmsales_nameemployee}</td>
                                <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmsales_nmwebagen}</td>
                                <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmsales_idwebagen}</td>
                                <td nowrap style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(rec.crmsales_deposit)}</td>
                            </tr>
                        {/each}
                        
                    </tbody>
                </table>
            </div>
            <div class="tab-pane" id="sales_noanswer">
                <table class="table table-sm">
                    <thead>
                        <tr>
                            <th nowrap width="7%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                            <th nowrap width="20%" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TEAM SALES</th>
                            <th nowrap width="*" style="text-align:left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NOTE</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each listcrmsales_noanswer as rec}
                            <tr>
                                <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmsales_status}</td>
                                <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmsales_nameemployee}</td>
                                <td nowrap style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.crmsales_note}</td>
                            </tr>
                        {/each}
                        
                    </tbody>
                </table>
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
	</slot:template>
</Modal>
