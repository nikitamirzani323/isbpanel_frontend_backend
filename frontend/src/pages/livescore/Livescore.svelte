<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listHomeDetail = [];
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
                page: "DOMAIN-VIEW",
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
        const res = await fetch("/api/livescore", {
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
                let record2 = record['category']
                let no = 0;
                totalrecord = 0
                for (var i = 0; i < record2.length; i++) {
                    no = no + 1;
                    totalrecord = totalrecord + 1;
                    let record_matches = record2[i]["matches"]["match"]
                    listHomeDetail = [];
                    if(record_matches.length != undefined){
                        for (var j = 0; j < record_matches.length; j++) {
                            console.log(
                                record_matches[j]["@status"]+" "+
                                record_matches[j]["@date"]+" "+
                                record_matches[j]["@formatted_date"]+" "+
                                record_matches[j]["localteam"]["@id"]+" "+
                                record_matches[j]["localteam"]["@name"]+" "+
                                record_matches[j]["localteam"]["@goals"]
                                +" VS "+
                                record_matches[j]["visitorteam"]["@id"]+" "+
                                record_matches[j]["visitorteam"]["@name"]+" "+
                                record_matches[j]["visitorteam"]["@goals"]
                            )
                            listHomeDetail = [
                                ...listHomeDetail,
                                {
                                    detail_status: record_matches[j]["@status"],
                                    detail_date: record_matches[j]["@date"],
                                    detail_formatedate: record_matches[j]["@formatted_date"],
                                    detail_localteam_id: record_matches[j]["localteam"]["@id"],
                                    detail_localteam_name: record_matches[j]["localteam"]["@name"],
                                    detail_localteam_goal: record_matches[j]["localteam"]["@goals"],
                                    detail_visitorteam_id: record_matches[j]["visitorteam"]["@id"],
                                    detail_visitorteam_name: record_matches[j]["visitorteam"]["@name"],
                                    detail_visitorteam_goal: record_matches[j]["visitorteam"]["@goals"],
                                },
                            ];
                        }
                    }
                    
                    
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_filegroup: record2[i]["@file_group"],
                            home_gid: record2[i]["@gid"],
                            home_id: record2[i]["@id"],
                            home_iscup: record2[i]["@iscup"],
                            home_name: record2[i]["@name"],
                            home_detail: listHomeDetail,
                        },
                    ];
                }
            }
        } else {
            // logout();
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