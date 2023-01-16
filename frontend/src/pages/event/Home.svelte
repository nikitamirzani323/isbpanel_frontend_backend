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
    let myModal_newentry = "";
    let idwebsite_field = 0;
    let nmevent_field = "";
    let startevent_field = "";
    let startevent_jam_field = "00:00";
    let endevent_field = "";
    let endevent_jam_field = "00:00";
    let create_field = "";
    let update_field = "";
    let listwebsiteagen_db = [];
    
    let username_agen_field = "";
    let searchMember = "";
    let filterMember = [];
    let css_loader = "display: none;";
    let msgloader = "";
    
   
    $: {
        if (searchMember) {
            filterMember = listHome.filter(
                (item) =>
                    item.home_phone
                        .toLowerCase()
                        .includes(searchMember.toLowerCase())
            );
        } else {
            filterMember = [...listHome];
        }
        
    }
    const NewData = (e,idwebsite,start,create,update) => {
        sData = e
        call_websiteagen()
        if(sData == "New"){
            clearField()
        }else{
            // alert(idwebsite)
            idwebsite_field = idwebsite;
            nmevent_field = "";
            startevent_field = dayjs(start).format("MM / DD / YYYY");
            startevent_jam_field = dayjs(start).format("HH:MM");
            endevent_field = "";
            endevent_jam_field = "00:00";
            create_field = create;
            update_field = update;
            
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
   
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
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
    
    function clearField(){
        idwebsite_field = 0;
        nmevent_field = "";
        startevent_field = "";
        startevent_jam_field = "00:00";
        endevent_field = "";
        endevent_jam_field = "00:00";
        create_field = "";
        update_field = "";
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
            case "SAVE_WEBSITEAGEN":
                handle_newWebsiteAgen();break;
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
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">WEBSITE AGEN</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAMA</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">START</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">END</th>
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
                                                NewData("Edit",rec.home_idwebsite, rec.home_start, rec.home_agen,rec.home_create,rec.home_update);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_websiteagen}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_start}</td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_end}</td>
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
    modal_body_css="height:400px;overflow-y: scroll;"
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



