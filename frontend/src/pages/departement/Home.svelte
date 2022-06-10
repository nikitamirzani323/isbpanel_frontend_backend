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
	let title_page = "DEPARTEMENT"
    let sData = "";
    let myModal_newentry = "";
    let departement_field = "";
    let idrecord = "";
    let flag_idrecord = false;
    let searchDepartement = "";
    let filterDepartement = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchDepartement) {
            filterDepartement = listHome.filter(
                (item) =>
                    item.departement_name
                        .toLowerCase()
                        .includes(searchDepartement.toLowerCase())
            );
        } else {
            filterDepartement = [...listHome];
        }
    }
    
    const NewData = (e,id,nm) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_idrecord = true
            idrecord = id
            departement_field = nm;
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
        if(idrecord == ""){
            flag = false
            msg += "The ID is required\n"
        }
        if(departement_field == ""){
            flag = false
            msg += "The Departement is required\n"
        }
        
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/departementsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"DEPARTEMENT-SAVE",
                    departement_id: idrecord,
                    departement_name: departement_field,
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
        flag_idrecord = ""
        departement_field = "";
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
    function uppercase(element) {
        function onInput(event) {
            element.value = element.value.toUpperCase();
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
                            bind:value={searchDepartement}
                            type="text"
                            class="form-control"
                            placeholder="Search Departement"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped table-hover">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPARTEMENT</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterDepartement as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewData("Edit",rec.departement_id,rec.departement_name);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.departement_no}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.departement_name}</td>
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
            <label for="exampleForm" class="form-label">ID</label>
            {#if flag_idrecord}
                <input
                    use:uppercase
                    bind:value={idrecord}
                    disabled 
                    class="required form-control"
                    type="text"
                    maxlength="5"
                    placeholder="ID"/>
            {:else}
                <input
                    use:uppercase
                    bind:value={idrecord} 
                    class="required form-control"
                    type="text"
                    maxlength="5"
                    placeholder="ID"/>
            {/if}
            
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Departement</label>
            <Input
                bind:value={departement_field}
                class="required"
                type="text"
                placeholder="Departement"/>
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



