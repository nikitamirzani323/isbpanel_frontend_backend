<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listDepartement = [];
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
                page: "EMPLOYEE-VIEW",
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
        const res = await fetch("/api/employee", {
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
            let record_depart = json.listdepartement;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                let employee_statuscss = ""
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["employee_status"] == "Y"){
                        employee_statuscss = "background:#FFEB3B;font-weight:bold;color:black;"
                    }else{
                        employee_statuscss = "background:#E91E63;font-weight:bold;color:white;"
                    }
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            employee_no: no,
                            employee_username: record[i]["employee_username"],
                            employee_iddepart: record[i]["employee_iddepart"],
                            employee_nmdepart: record[i]["employee_nmdepart"],
                            employee_name: record[i]["employee_name"],
                            employee_phone: record[i]["employee_phone"],
                            employee_status: record[i]["employee_status"],
                            employee_statuscss: employee_statuscss,
                            employee_create: record[i]["employee_create"],
                            employee_update: record[i]["employee_update"],
                        },
                    ];
                }
            }
            for (var i = 0; i < record_depart.length; i++) {
                listDepartement = [
                    ...listDepartement,
                    {
                        departement_id: record_depart[i]["departement_id"],
                        departement_name: record_depart[i]["departement_name"],
                    },
                ];
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
        listDepartement = [];
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
    {listDepartement}
    {totalrecord}/>
{/if}