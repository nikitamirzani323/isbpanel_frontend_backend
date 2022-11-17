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
	export let listDepartement = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "EMPLOYEE"
    let sData = "";
    let myModal_newentry = "";
    let password_field = "";
    let iddepart_field = "";
    let name_field = "";
    let phone_field = "";
    let status_field = "";
    let idrecord = "";
    let flag_idrecord = false;
    let searchEmployee = "";
    let filterEmployee = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchEmployee) {
            filterEmployee = listHome.filter(
                (item) =>
                    item.employee_username
                        .toLowerCase()
                        .includes(searchEmployee.toLowerCase()) || 
                    item.employee_nmdepart
                        .toLowerCase()
                        .includes(searchEmployee.toLowerCase())
            );
        } else {
            filterEmployee = [...listHome];
        }
    }
    
    const NewData = (e,username,iddepart,name,phone,status) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_idrecord = true
            idrecord = username
            iddepart_field = iddepart;
            name_field = name;
            phone_field = phone;
            status_field = status;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData=="New"){
            if(password_field == ""){
                flag = false
                msg += "The Password is required\n"
            }
        }
        if(idrecord == ""){
            flag = false
            msg += "The ID is required\n"
        }
        if(iddepart_field == ""){
            flag = false
            msg += "The Departement is required\n"
        }
        if(name_field == ""){
            flag = false
            msg += "The Nama is required\n"
        }
        if(phone_field == ""){
            flag = false
            msg += "The Phone is required\n"
        }
        if(status_field == ""){
            flag = false
            msg += "The Status is required\n"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/employeesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"EMPLOYEE-SAVE",
                    employee_username: idrecord,
                    employee_password: password_field,
                    employee_iddepart: iddepart_field,
                    employee_name: name_field,
                    employee_phone: phone_field,
                    employee_status: status_field,
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
        flag_idrecord = false
        password_field = "";
        iddepart_field = "";
        name_field = "";
        phone_field = "";
        status_field = "";
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
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
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
                            bind:value={searchEmployee}
                            type="text"
                            class="form-control"
                            placeholder="Search Employee"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">STATUS</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPARTEMENT</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterEmployee as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewData("Edit",rec.employee_username,rec.employee_iddepart,rec.employee_name,rec.employee_phone,rec.employee_status);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.employee_no}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.employee_statuscss}">
                                            {status(rec.employee_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.employee_nmdepart}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.employee_username}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.employee_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.employee_phone}</td>
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
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Username</label>
            {#if flag_idrecord}
                <input
                    use:lowercase
                    bind:value={idrecord}
                    disabled 
                    class="required form-control"
                    type="text"
                    maxlength="10"
                    placeholder="Username"/>
            {:else}
                <input
                    use:lowercase
                    bind:value={idrecord} 
                    class="required form-control"
                    type="text"
                    maxlength="10"
                    placeholder="Username"/>
            {/if}
            
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Departement</label>
            <select
                bind:value="{iddepart_field}" 
                name="departement" id="departement" 
                class="required form-control">
                {#each listDepartement as rec}
                <option value="{rec.departement_id}">{rec.departement_name}</option>
                {/each}
            </select>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Password</label>
            <input
                bind:value={password_field}
                type="password"
                class="form-control required"
                placeholder="Password"
                aria-label="Password"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Nama</label>
            <Input
                bind:value={name_field}
                class="required"
                type="text"
                placeholder="Nama"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Phone</label>
            <Input
                bind:value={phone_field}
                class="required"
                type="text"
                placeholder="Phone"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select
                class="form-control required"
                bind:value={status_field}>
                <option value="Y">ACTIVE</option>
                <option value="N">DEACTIVE</option>
            </select>
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



