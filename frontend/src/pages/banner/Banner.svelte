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
                page: "BANNER-VIEW",
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
        const res = await fetch("/api/banner", {
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
                let banner_statuscss = ""
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["banner_status"] == "Y"){
                        banner_statuscss = "background:#FFEB3B;font-weight:bold;color:black;"
                    }else{
                        banner_statuscss = "background:#E91E63;font-weight:bold;color:white;"
                    }
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            banner_no: no,
                            banner_id: record[i]["banner_id"],
                            banner_name: record[i]["banner_name"],
                            banner_url: record[i]["banner_url"],
                            banner_urlwebsite: record[i]["banner_urlwebsite"],
                            banner_posisi: record[i]["banner_posisi"],
                            banner_device: record[i]["banner_device"],
                            banner_display: record[i]["banner_display"],
                            banner_status: record[i]["banner_status"],
                            banner_statuscss: banner_statuscss,
                            banner_create: record[i]["banner_create"],
                            banner_update: record[i]["banner_update"],
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