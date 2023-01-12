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
	let title_page = "MEMBER"
    let sData = "";
    let myModal_newentry = "";
    let phone_flag = false;
    let phone_field = "";
    let nama_field = "";
    let create_field = "";
    let update_field = "";
    let listwebsiteagen_db = [];
    let list_agen = [];
    let list_agen_count = 0;
    let idwebsite_agen_field = 0;
    let nmwebsite_agen_field = "";
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
    const NewData = (e,phone,nama,dataagen,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            list_agen = [];
            list_agen_count = 0;
            phone_flag = true;
            phone_field = phone;
            nama_field = nama;
            create_field = create;
            update_field = update;
            if(dataagen != null){
                for (var i = 0; i < dataagen.length; i++) {
                    list_agen_count = list_agen_count + 1;
                    list_agen = [
                        ...list_agen,
                        {
                            agen_no: list_agen_count,
                            agen_idwebsite: dataagen[i].memberagen_idwebagen,
                            agen_nmwebsite: dataagen[i].memberagen_website,
                            agen_username: dataagen[i].memberagen_username,
                        },
                    ];
                }
            }
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const ShowFormAgen = () => {
        call_websiteagen()
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalformagen"));
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
    async function handle_newWebsiteAgen() {
        if (username_agen_field != "" || idwebsite_agen_field != "") {
            list_agen_count = list_agen_count + 1;
            let obj = listwebsiteagen_db.find(o => o.websiteagen_id === idwebsite_agen_field);
            nmwebsite_agen_field = obj.websiteagen_name;
            list_agen = [
                ...list_agen,
                {
                    agen_no: parseInt(list_agen_count),
                    agen_idwebsite: idwebsite_agen_field,
                    agen_nmwebsite: nmwebsite_agen_field,
                    agen_username: username_agen_field,
                },
            ];
        } else {
            alert("The Website Agen + Username is required");
        }
        idwebsite_agen_field = 0;
        nmwebsite_agen_field = ""
        username_agen_field = "";
    }
    async function handle_deleteWebsiteAgen(e) {
        let temp = list_agen.filter(
            (item) => item.agen_no !== parseInt(e)
        );
        list_agen = [];
        for (var i = 0; i < temp.length; i++) {
            list_agen = [
                ...list_agen,
                {
                    agen_no: parseInt(temp[i].agen_no),
                    agen_idwebsite: temp[i].agen_idwebsite,
                    agen_nmwebsite: temp[i].agen_nmwebsite,
                    agen_username: temp[i].agen_username,
                },
            ];
        }
    }
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(nama_field == ""){
                flag = false
                msg += "The Nama is required\n"
            }
            if(list_agen.length < 1){
                flag = false
                msg += "The Website Agen is required\n"
            }
        }else{
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(nama_field == ""){
                flag = false
                msg += "The Nama is required\n"
            }
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/membersave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"DOMAIN-SAVE",
                    member_phone: phone_field,
                    member_name: nama_field,
                    member_agen: list_agen,
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
        list_agen = [];
        list_agen_count = 0;
        phone_flag = false;
        phone_field = "";
        nama_field = "";
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
                            placeholder="Search Member"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="8%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAMA</th>
                                <th NOWRAP width="25%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">WEBSITE AGEN</th>
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
                                                NewData("Edit",rec.home_phone, rec.home_name, rec.home_agen,rec.home_create,rec.home_update);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_phone}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {#if rec.home_agen != null}
                                        {#each rec.home_agen as rec2}
                                        {rec2.memberagen_website} - {rec2.memberagen_username}<br />
                                        {/each}
                                        {/if}
                                    </td>
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
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_body_css="height:250px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone</label>
                    <Input
                        bind:value={phone_field}
                        disabled='{phone_flag}'
                        class="required"
                        type="text"
                        placeholder="phone"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Nama</label>
                    <Input
                        bind:value={nama_field}
                        class="required"
                        type="text"
                        placeholder="nama"/>
                </div>
                {#if sData == "Edit"}
                    <div class="alert alert-dark" role="alert" style="font-size:11px;padding:5px;">
                        Create : {create_field}
                        <br />
                        Update : {update_field}
                    </div>
                {/if}
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                  <label for="exampleForm" class="form-label">Agen -
                    <span
                      on:click={() => {
                        ShowFormAgen();
                      }} style="text-decoration: underline;cursor:pointer;color:blue;">New</span>
                  </label>
                  <table class="table table-sm">
                    <tbody>
                      {#each list_agen as rec}
                        <tr>
                          <td width="1%" style="cursor: pointer;">
                            <i on:click={() => { handle_deleteWebsiteAgen(rec.agen_no);}} class="bi bi-trash"/>
                          </td>
                          <td width="*" style="text-align:left;vertical-align:top;font-size:12px;color:blue;">
                            {rec.agen_nmwebsite} - {rec.agen_username}
                          </td>
                        </tr>
                      {/each}
                    </tbody>
                  </table>
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

<Modal
  modal_id="modalformagen"
  modal_size="modal-dialog-centered"
  modal_title="Agen"
  modal_body_css=""
  modal_footer_css="padding:5px;"
  modal_footer={true}>
  <slot:template slot="body">
    <div class="mb-3">
        <label for="exampleForm" class="form-label">Website</label>
        <select bind:value={idwebsite_agen_field} class="form-control required">
            {#each listwebsiteagen_db as rec}
                <option value="{rec.websiteagen_id}">{rec.websiteagen_name}</option>
            {/each}
        </select>
      </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Username</label>
      <Input
        bind:value={username_agen_field}
        class="required"
        type="text"
        placeholder="Username"/>
    </div>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="SAVE_WEBSITEAGEN"
      button_title="Save"
      button_css="btn-warning"/>
  </slot:template>
</Modal>

