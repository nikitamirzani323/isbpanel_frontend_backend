<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let record = "";
    let record_message = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "SLOT-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch("/api/providerslot", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                let providerslot_statuscss = ""
                let providerslot_display_css = ""
                let providerslot_counter_css = ""
                let providerslot_totalgameslot_css = ""
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(record[i]["providerslot_status"] == "Y"){
                        providerslot_statuscss = "background:#FFEB3B;font-weight:bold;color:black;"
                    }else{
                        providerslot_statuscss = "background:#E91E63;font-weight:bold;color:white;"
                    }
                    if(parseInt(record[i]["providerslot_display"])>0){
                        providerslot_display_css = "color:blue;font-weight:bold;"
                    }else{
                        providerslot_display_css = "color:red;font-weight:bold;"
                    }
                    if(parseInt(record[i]["providerslot_counter"])>0){
                        providerslot_counter_css = "color:blue;font-weight:bold;"
                    }else{
                        providerslot_counter_css = "color:red;font-weight:bold;"
                    }
                    if(parseInt(record[i]["providerslot_totalgameslot"])>0){
                        providerslot_totalgameslot_css = "color:blue;font-weight:bold;"
                    }else{
                        providerslot_totalgameslot_css = "color:red;font-weight:bold;"
                    }
                    listHome = [
                        ...listHome,
                        {
                            providerslot_no: no,
                            providerslot_id: record[i]["providerslot_id"],
                            providerslot_name: record[i]["providerslot_name"],
                            providerslot_display: record[i]["providerslot_display"],
                            providerslot_display_css: providerslot_display_css,
                            providerslot_counter: record[i]["providerslot_counter"],
                            providerslot_counter_css: providerslot_counter_css,
                            providerslot_totalgameslot: record[i]["providerslot_totalgameslot"],
                            providerslot_totalgameslot_css: providerslot_totalgameslot_css,
                            providerslot_image: record[i]["providerslot_image"],
                            providerslot_slug: record[i]["providerslot_slug"],
                            providerslot_title: record[i]["providerslot_title"],
                            providerslot_descp: record[i]["providerslot_descp"],
                            providerslot_status: record[i]["providerslot_status"],
                            providerslot_statuscss: providerslot_statuscss,
                            providerslot_create: record[i]["providerslot_create"],
                            providerslot_update: record[i]["providerslot_update"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {totalrecord}/>
{/if}