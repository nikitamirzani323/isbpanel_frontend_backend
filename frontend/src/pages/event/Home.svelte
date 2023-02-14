<script>
    import { Input } from "sveltestrap";
    import { initializeApp } from "firebase/app";
    import { getDatabase, ref, set } from "firebase/database";
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
    let sDataNewPartisipasi = "";
    let myModal_newentry = "";
    let myModal_memberagen = "";
    let myModal_memberagenwinner = "";
    const firebaseConfig = {
        apiKey: "AIzaSyDgzVaNSmYg80ycxn2fA_Rsbgji2jUbqf8",
        authDomain: "isbproject-5028a.firebaseapp.com",
        databaseURL: "https://isbproject-5028a-default-rtdb.asia-southeast1.firebasedatabase.app",
        projectId: "isbproject-5028a",
        storageBucket: "isbproject-5028a.appspot.com",
        messagingSenderId: "181461320518",
        appId: "1:181461320518:web:a3280cd5dde31b9ebe5b47",
        measurementId: "G-RYBN4R3HVH"
    };
    const app = initializeApp(firebaseConfig);
    const db = getDatabase(app);

    let idevent_global = 0;
    let idwebsite_global = 0;
    let nmwebsite_global = "";
    let nmwevent_global = "";
    let deposit_global = 0;
    let username_global = "";
    let fieldwinner_global = "";
    let fieldwinnerusername_global = "";

    let idwebsite_field = 0;
    let nmevent_field = "";
    let startevent_field = "";
    let startevent_jam_field = "00:00";
    let endevent_field = "";
    let endevent_jam_field = "00:00";
    let mindeposit_field = 0;
    let flag_buttonsave = true;
    let create_field = "";
    let update_field = "";

    let listmembergroup_total = 0;
    let listpartisipasi_total = 0;
    let listpartisipasivoucher_total = 0;
    let listpartisipasi_db = [];
    let listpartisipasivoucher_db = [];
    let listpartisipasivoucherwinner_db = [];
    let listmembergroup_db = [];
    let listwebsiteagen_db = [];
    let listmemberagen_db = [];
    let listfirebase_db = [];

    
    let idmemberagen_partisipasi_field = 0;
    let username_partisipasi_field = "";
    let qty_partisipasi_field = 0;

    let prize1_username  = "";
    let prize1_winner  = "";
    let prize1_winner_flag  = false;
    let prize1_winner_save_flag  = false;
    let prize2_username = "";
    let prize2_winner = "";
    let prize2_winner_flag  = false;
    let prize2_winner_save_flag  = false;
    let prize3_username = "";
    let prize3_winner = "";
    let prize3_winner_flag  = false;
    let prize3_winner_save_flag  = false;
    let prize4_username = "";
    let prize4_winner = "";
    let prize4_winner_flag  = false;
    let prize4_winner_save_flag  = false;
    let prize5_username = "";
    let prize5_winner = "";
    let prize5_winner_flag  = false;
    let prize5_winner_save_flag  = false;
    let prize6_username = "";
    let prize6_winner = "";
    let prize6_winner_flag  = false;
    let prize6_winner_save_flag  = false;
    let prize7_username = "";
    let prize7_winner = "";
    let prize7_winner_flag  = false;
    let prize7_winner_save_flag  = false;
    let prize8_username = "";
    let prize8_winner = "";
    let prize8_winner_flag  = false;
    let prize8_winner_save_flag  = false;
    let prize9_username = "";
    let prize9_winner = "";
    let prize9_winner_flag  = false;
    let prize9_winner_save_flag  = false;
    let prize10_username = "";
    let prize10_winner = "";
    let prize10_winner_flag  = false;
    let prize10_winner_save_flag  = false;
    let prize11_username = "";
    let prize11_winner = "";
    let prize11_winner_flag  = false;
    let prize11_winner_save_flag  = false;
    let prize12_username = "";
    let prize12_winner = "";
    let prize12_winner_flag  = false;
    let prize12_winner_save_flag  = false;
    let prize13_username = "";
    let prize13_winner = "";
    let prize13_winner_flag  = false;
    let prize13_winner_save_flag  = false;
    let prize14_username = "";
    let prize14_winner = "";
    let prize14_winner_flag  = false;
    let prize14_winner_save_flag  = false;
    let prize15_username = "";
    let prize15_winner = "";
    let prize15_winner_flag  = false;
    let prize15_winner_save_flag  = false;
    let prize16_username = "";
    let prize16_winner = "";
    let prize16_winner_flag  = false;
    let prize16_winner_save_flag  = false;
    let prize17_username = "";
    let prize17_winner = "";
    let prize17_winner_flag  = false;
    let prize17_winner_save_flag  = false;
    let prize18_username = "";
    let prize18_winner = "";
    let prize18_winner_flag  = false;
    let prize18_winner_save_flag  = false;
    let prize19_username = "";
    let prize19_winner = "";
    let prize19_winner_flag  = false;
    let prize19_winner_save_flag  = false;
    let prize20_username = "";
    let prize20_winner = "";
    let prize20_winner_flag  = false;
    let prize20_winner_save_flag  = false;

    let panel_footer_listall = true;
    let panel_footer_member = false;

    let listvoucher_view = [];
    let searchListpartisipasi = "";
    let filterListpartisipasi = [];
    let searchMember = "";
    let filterMember = [];
    let css_loader = "display: none;";
    let msgloader = "";
    
   
    $: {
        if (searchMember) {
            filterMember = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchMember.toLowerCase())
            );
        } else {
            filterMember = [...listHome];
        }
        
        if (searchListpartisipasi) {
            filterListpartisipasi = listpartisipasi_db.filter(
                (item) =>
                    item.eventdetail_voucher
                        .toLowerCase()
                        .includes(searchListpartisipasi.toLowerCase()) ||
                    item.eventdetail_username
                        .toLowerCase()
                        .includes(searchListpartisipasi.toLowerCase())
            );
        } else {
            filterListpartisipasi = [...listpartisipasi_db];
        }
    }
    const NewData = (e,idevent, idwebsite,event,start,end,deposit, status, create,update) => {
        sData = e
        call_websiteagen()
        if(sData == "New"){
            clearField()
        }else{
            if(status == "OFFLINE"){
                flag_buttonsave = false;
            }else{
                flag_buttonsave = true;
            }
            idevent_global = idevent
            idwebsite_field = idwebsite;
            nmevent_field = event;
            startevent_field = dayjs(start).format("YYYY-MM-DD");
            startevent_jam_field = dayjs(start).format("HH:MM");
            endevent_field = dayjs(end).format("YYYY-MM-DD");
            endevent_jam_field = dayjs(end).format("HH:MM");
            mindeposit_field = deposit;
            create_field = create;
            update_field = update;
            
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const ListPartisipasi = (idevent,idwebsite,nmwebsite,nmevent,deposit,status) => {
        call_listpartisipasi(idevent,0)
        if(status == "OFFLINE"){
            flag_buttonsave = false;
        }else{
            flag_buttonsave = true;
        }
        idevent_global = idevent
        idwebsite_global = idwebsite
        nmwebsite_global = nmwebsite
        nmwevent_global = nmevent
        deposit_global = deposit
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistpartisipasi"));
        myModal_newentry.show();
    };
    const NewFormPartisipasi = () => {
        sDataNewPartisipasi = "New";
        clearFieldPartisipasi();
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrypartisipasi"));
        myModal_newentry.show();
    };
    const ListMemberAgen = () => {
        call_memberagen(idwebsite_global)
        myModal_memberagen = new bootstrap.Modal(document.getElementById("modallistmemberagen"));
        myModal_memberagen.show();
    };
    const ListMemberAgenVoucher = (e,f,tipe) => {
        username_global = f;
        call_listpartisipasi(idevent_global,e)   
        if(tipe == ""){
            myModal_memberagen = new bootstrap.Modal(document.getElementById("modallistmemberagenvoucher"));
            myModal_memberagen.show();
        }else{
            fieldwinnerusername_global = f;
            myModal_memberagenwinner = new bootstrap.Modal(document.getElementById("modallistmemberagenvoucherwinner"));
            myModal_memberagenwinner.show();
        }
    };
  
  
    const TabPartisipasi = (e) => {
        switch (e){
            case "LISTALL":
                panel_footer_listall = true;
                panel_footer_member = false;
                call_listpartisipasi(idevent_global,0);break;
            case "MEMBER":
                panel_footer_member = true;
                panel_footer_listall = false;
                call_listpartisipasimembergroup(idevent_global);break;
            case "WINNER":
                mappingwinner(idevent_global);
                panel_footer_member = false;
                panel_footer_listall = false;break;
        }
    };
    const InsertPartisipasi = (id,e) => {
        idmemberagen_partisipasi_field = parseInt(id);
        username_partisipasi_field = e;
        myModal_memberagen.hide();
    };
    const InsertFirebase = (e,f) => {
        let flag = true;
        let ideventdetail = f
        let idevent = idevent_global
        let winnerposition = fieldwinner_global
        if(listfirebase_db.length > 0){
            for(let i=0;i<listfirebase_db.length;i++){
                if(listfirebase_db[i]==e){
                    flag = false;
                }
            }
        }
        if(flag){
            handleSaveWinner(idevent,ideventdetail,winnerposition)
            listfirebase_db.push(e)
            switch (fieldwinner_global){
                case "PRIZE_1":
                    prize1_username = fieldwinnerusername_global;
                    prize1_winner = e;
                    prize1_winner_flag = true;
                    prize1_winner_save_flag = true;
                    break;
                case "PRIZE_2":
                    prize2_username = fieldwinnerusername_global;
                    prize2_winner = e;
                    prize2_winner_flag = true;
                    prize2_winner_save_flag = true;
                    break;
                case "PRIZE_3":
                    prize3_username = fieldwinnerusername_global;
                    prize3_winner = e;
                    prize3_winner_flag = true;
                    prize3_winner_save_flag = true;
                    break;
                case "PRIZE_4":
                    prize4_username = fieldwinnerusername_global;
                    prize4_winner = e;
                    prize4_winner_flag = true;
                    prize4_winner_save_flag = true;
                    break;
                case "PRIZE_5":
                    prize5_username = fieldwinnerusername_global;
                    prize5_winner = e;
                    prize5_winner_flag = true;
                    prize5_winner_save_flag = true;
                    break;
                case "PRIZE_6":
                    prize6_username = fieldwinnerusername_global;
                    prize6_winner = e;
                    prize6_winner_flag = true;
                    prize6_winner_save_flag = true;
                    break;
                case "PRIZE_7":
                    prize7_username = fieldwinnerusername_global;
                    prize7_winner = e;
                    prize7_winner_flag = true;
                    prize7_winner_save_flag = true;
                    break;
                case "PRIZE_8":
                    prize8_username = fieldwinnerusername_global;
                    prize8_winner = e;
                    prize8_winner_flag = true;
                    prize8_winner_save_flag = true;
                    break;
                case "PRIZE_9":
                    prize9_username = fieldwinnerusername_global;
                    prize9_winner = e;
                    prize9_winner_flag = true;
                    prize9_winner_save_flag = true;
                    break;
                case "PRIZE_10":
                    prize10_username = fieldwinnerusername_global;
                    prize10_winner = e;
                    prize10_winner_flag = true;
                    prize10_winner_save_flag = true;
                    break;
                case "PRIZE_11":
                    prize11_username = fieldwinnerusername_global;
                    prize11_winner = e;
                    prize11_winner_flag = true;
                    prize11_winner_save_flag = true;
                    break;
                case "PRIZE_12":
                    prize12_username = fieldwinnerusername_global;
                    prize12_winner = e;
                    prize12_winner_flag = true;
                    prize12_winner_save_flag = true;
                    break;
                case "PRIZE_13":
                    prize13_username = fieldwinnerusername_global;
                    prize13_winner = e;
                    prize13_winner_flag = true;
                    prize13_winner_save_flag = true;
                    break;
                case "PRIZE_14":
                    prize14_username = fieldwinnerusername_global;
                    prize14_winner = e;
                    prize14_winner_flag = true;
                    prize14_winner_save_flag = true;
                    break;
                case "PRIZE_15":
                    prize15_username = fieldwinnerusername_global;
                    prize15_winner = e;
                    prize15_winner_flag = true;
                    prize15_winner_save_flag = true;
                    break;
                case "PRIZE_16":
                    prize16_username = fieldwinnerusername_global;
                    prize16_winner = e;                
                    prize16_winner_flag = true;
                    prize16_winner_save_flag = true;
                    break;
                case "PRIZE_17":
                    prize17_username = fieldwinnerusername_global;
                    prize17_winner = e;
                    prize17_winner_flag = true;
                    prize17_winner_save_flag = true;
                    break;
                case "PRIZE_18":
                    prize18_username = fieldwinnerusername_global;
                    prize18_winner = e;
                    prize18_winner_flag = true;
                    prize18_winner_save_flag = true;
                    break;
                case "PRIZE_19":
                    prize19_username = fieldwinnerusername_global;
                    prize19_winner = e;
                    prize19_winner_flag = true;
                    prize19_winner_save_flag = true;
                    break;
                case "PRIZE_20":
                    prize20_username = fieldwinnerusername_global;
                    prize20_winner = e;
                    prize20_winner_flag = true;
                    prize20_winner_save_flag = true;
                    break;
            }
            set(ref(db, 'eventisbproject'), {
                prize1: prize1_winner,
                prize2: prize2_winner,
                prize3: prize3_winner,
                prize4: prize4_winner,
                prize5: prize5_winner,
                prize6: prize6_winner,
                prize7: prize7_winner,
                prize8: prize8_winner,
                prize9: prize9_winner,
                prize10: prize10_winner,
                prize11: prize11_winner,
                prize12: prize12_winner,
                prize13: prize13_winner,
                prize14: prize14_winner,
                prize15: prize15_winner,
                prize16: prize16_winner,
                prize17: prize17_winner,
                prize18: prize18_winner,
                prize19: prize19_winner,
                prize20: prize20_winner,
            });
            myModal_memberagen.hide();
            myModal_memberagenwinner.hide();
        }else{
            alert("Duplicate Voucher")
        }
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function call_listpartisipasi(idevent,idmember) {
        if(idmember > 0){     
            listpartisipasivoucher_db = [];
            listpartisipasivoucher_total = 0;
        }else{
            listpartisipasi_db = [];
            listpartisipasi_total = 0;
        }
        const res = await fetch("/api/eventdetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "MOVIEALBUM-VIEW",
                event_id: idevent,
                event_idmemberagen: parseInt(idmember),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(idmember > 0){
                        listpartisipasivoucher_total = listpartisipasivoucher_total + record[i]["eventdetail_deposit"];
                        listpartisipasivoucher_db = [
                            ...listpartisipasivoucher_db,
                            {
                                eventdetail_no: no,
                                eventdetail_id: record[i]["eventdetail_id"],
                                eventdetail_phone: record[i]["eventdetail_phone"],
                                eventdetail_username: record[i]["eventdetail_username"],
                                eventdetail_voucher: record[i]["eventdetail_voucher"],
                                eventdetail_deposit: record[i]["eventdetail_deposit"],
                                eventdetail_status: record[i]["eventdetail_status"],
                                eventdetail_create: record[i]["eventdetail_create"],
                                eventdetail_update: record[i]["eventdetail_update"],
                            },
                        ];
                    }else{
                        listpartisipasi_total = listpartisipasi_total + record[i]["eventdetail_deposit"];
                        listpartisipasi_db = [
                            ...listpartisipasi_db,
                            {
                                eventdetail_no: no,
                                eventdetail_id: record[i]["eventdetail_id"],
                                eventdetail_phone: record[i]["eventdetail_phone"],
                                eventdetail_username: record[i]["eventdetail_username"],
                                eventdetail_voucher: record[i]["eventdetail_voucher"],
                                eventdetail_deposit: record[i]["eventdetail_deposit"],
                                eventdetail_status: record[i]["eventdetail_status"],
                                eventdetail_create: record[i]["eventdetail_create"],
                                eventdetail_update: record[i]["eventdetail_update"],
                            },
                        ];
                    }
                }
            }
        }
    }
    async function call_listpartisipasimembergroup(idevent) {
        listmembergroup_db = [];
        listmembergroup_total = 0;
        const res = await fetch("/api/eventgroupdetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "MOVIEALBUM-VIEW",
                event_id: idevent,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listmembergroup_total = listmembergroup_total + record[i]["eventdetailgroup_deposit"];
                    listmembergroup_db = [
                        ...listmembergroup_db,
                        {
                            eventdetailgroup_no: no,
                            eventdetailgroup_idmember: record[i]["eventdetailgroup_idmember"],
                            eventdetailgroup_phone: record[i]["eventdetailgroup_phone"],
                            eventdetailgroup_username: record[i]["eventdetailgroup_username"],
                            eventdetailgroup_deposit: record[i]["eventdetailgroup_deposit"],
                            eventdetailgroup_voucher: record[i]["eventdetailgroup_voucher"],
                        },
                    ];
                }
            }
        }
    }
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
    async function call_memberagen(e) {
        listmemberagen_db = [];
        const res = await fetch("/api/memberselect", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "MOVIEALBUM-VIEW",
                memberagen_idwebagen: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listmemberagen_db = [
                        ...listmemberagen_db,
                        {
                            memberagen_no: no,
                            memberagen_id: record[i]["memberagen_id"],
                            memberagen_username: record[i]["memberagen_username"],
                            memberagen_phone: record[i]["memberagen_phone"],
                            memberagen_name: record[i]["memberagen_name"],
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
            if(mindeposit_field == 0){
                flag = false
                msg += "The Minimal Deposit is required\n"
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
            if(mindeposit_field == 0){
                flag = false
                msg += "The Minimal Deposit is required\n"
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
                    event_id: idevent_global,
                    event_idwebagen: idwebsite_field,
                    event_name: nmevent_field,
                    event_mindeposit: parseInt(mindeposit_field),
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
    async function handleSavePartisipasi() {
        let flag = true
        let msg = ""
        if(sDataNewPartisipasi == "New"){
            if(idmemberagen_partisipasi_field == 0){
                flag = false
                msg += "The Member is required\n"
            }
            if(qty_partisipasi_field == ""){
                flag = false
                msg += "The Qty is required\n"
            }
        }else{
            if(idmemberagen_partisipasi_field == 0){
                flag = false
                msg += "The Member is required\n"
            }
            if(qty_partisipasi_field == ""){
                flag = false
                msg += "The Qty is required\n"
            }
        }
        
        if(flag){
            listvoucher_view = [];
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/eventdetailsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataNewPartisipasi,
                    page:"DOMAIN-SAVE",
                    eventdetail_id: idwebsite_field,
                    eventdetail_idevent: parseInt(idevent_global),
                    eventdetail_idmemberagen: parseInt(idmemberagen_partisipasi_field),
                    eventdetail_qty: parseInt(qty_partisipasi_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                let record = json.record;
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        listvoucher_view = [
                            ...listvoucher_view,
                            {
                                event_voucher: record[i]["event_voucher"],
                            },
                        ];
                    }
                    myModal_newentry = new bootstrap.Modal(document.getElementById("modalviewvoucher"));
                    myModal_newentry.show();
                }
                msgloader = json.message;
                call_listpartisipasi(idevent_global)
                clearFieldPartisipasi();
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
    async function handleSaveWinner(idevent,ideventdetail,winner) {
        let flag = true
        let msg = ""
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/eventdetailstatussave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: 'New',
                    page:"DOMAIN-SAVE",
                    eventdetail_id: ideventdetail,
                    eventdetail_idevent: parseInt(idevent),
                    eventdetail_status: winner,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                mappingwinner(idevent);
                msgloader = json.message;
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
    async function mappingwinner(idevent) {
        listpartisipasivoucherwinner_db = [];
        const res = await fetch("/api/eventdetailwinner", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: 'New',
                page: "MOVIEALBUM-VIEW",
                event_id: idevent,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    
                    listpartisipasivoucherwinner_db = [
                            ...listpartisipasivoucherwinner_db,
                            {
                                eventdetail_no: no,
                                eventdetail_id: record[i]["eventdetail_id"],
                                eventdetail_phone: record[i]["eventdetail_phone"],
                                eventdetail_username: record[i]["eventdetail_username"],
                                eventdetail_voucher: record[i]["eventdetail_voucher"],
                                eventdetail_deposit: record[i]["eventdetail_deposit"],
                                eventdetail_status: record[i]["eventdetail_status"],
                                eventdetail_create: record[i]["eventdetail_create"],
                                eventdetail_update: record[i]["eventdetail_update"],
                            },
                        ];
                }
                for(let i=0;i<listpartisipasivoucherwinner_db.length;i++){
                    switch(listpartisipasivoucherwinner_db[i].eventdetail_status){
                        case "PRIZE_1":
                            prize1_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize1_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize1_winner_flag  = true;
                            prize1_winner_save_flag  = true;
                            break;
                        case "PRIZE_2":
                            prize2_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize2_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize2_winner_flag  = true;
                            prize2_winner_save_flag  = true;
                            break;
                        case "PRIZE_3":
                            prize3_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize3_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize3_winner_flag  = true;
                            prize3_winner_save_flag  = true;
                            break;
                        case "PRIZE_4":
                            prize4_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize4_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize4_winner_flag  = true;
                            prize4_winner_save_flag  = true;
                            break;
                        case "PRIZE_5":
                            prize5_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize5_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize5_winner_flag  = true;
                            prize5_winner_save_flag  = true;
                            break;
                        case "PRIZE_6":
                            prize6_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize6_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize6_winner_flag  = true;
                            prize6_winner_save_flag  = true;
                            break;
                        case "PRIZE_7":
                            prize7_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize7_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize7_winner_flag  = true;
                            prize7_winner_save_flag  = true;
                            break;
                        case "PRIZE_8":
                            prize8_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize8_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize8_winner_flag  = true;
                            prize8_winner_save_flag  = true;
                            break;
                        case "PRIZE_9":
                            prize9_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize9_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize9_winner_flag  = true;
                            prize9_winner_save_flag  = true;
                            break;
                        case "PRIZE_10":
                            prize10_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize10_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize10_winner_flag  = true;
                            prize10_winner_save_flag  = true;
                            break;
                        case "PRIZE_11":
                            prize11_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize11_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize11_winner_flag  = true;
                            prize11_winner_save_flag  = true;
                            break;
                        case "PRIZE_12":
                            prize12_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize12_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize12_winner_flag  = true;
                            prize12_winner_save_flag  = true;
                            break;
                        case "PRIZE_13":
                            prize13_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize13_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize13_winner_flag  = true;
                            prize13_winner_save_flag  = true;
                            break;
                        case "PRIZE_14":
                            prize14_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize14_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize14_winner_flag  = true;
                            prize14_winner_save_flag  = true;
                            break;
                        case "PRIZE_15":
                            prize15_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize15_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize15_winner_flag  = true;
                            prize15_winner_save_flag  = true;
                            break;
                        case "PRIZE_16":
                            prize16_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize16_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize16_winner_flag  = true;
                            prize16_winner_save_flag  = true;
                            break;
                        case "PRIZE_17":
                            prize17_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize17_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize17_winner_flag  = true;
                            prize17_winner_save_flag  = true;
                            break;
                        case "PRIZE_18":
                            prize18_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize18_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize18_winner_flag  = true;
                            prize18_winner_save_flag  = true;
                            break;
                        case "PRIZE_19":
                            prize19_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize19_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize19_winner_flag  = true;
                            prize19_winner_save_flag  = true;
                            break;
                        case "PRIZE_20":
                            prize20_username = listpartisipasivoucherwinner_db[i].eventdetail_username;
                            prize20_winner  = listpartisipasivoucherwinner_db[i].eventdetail_voucher;
                            prize20_winner_flag  = true;
                            prize20_winner_save_flag  = true;
                            break;
                    }
                }
            }
        }
    }
    
    function handleMemberPopup(e){
        fieldwinner_global = e;
        call_listpartisipasimembergroup(idevent_global)
        
        myModal_memberagen = new bootstrap.Modal(document.getElementById("modallistmembervoucher"));
        myModal_memberagen.show();
    }
    function clearField(){
        idevent_global = 0;
        idwebsite_global = 0;
        nmwebsite_global = "";
        nmwevent_global = "";
        fieldwinner_global = "";
        idwebsite_field = 0;
        nmevent_field = "";
        startevent_field = "";
        startevent_jam_field = "00:00";
        endevent_field = "";
        endevent_jam_field = "00:00";
        mindeposit_field = 0;
        flag_buttonsave = true;
        create_field = "";
        update_field = "";
    }
    function clearFieldPartisipasi(){
        idmemberagen_partisipasi_field = 0;
        username_partisipasi_field = "";
        qty_partisipasi_field = 0;
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
            case "FORM_PARTISIPASI":
                NewFormPartisipasi();break;
            case "SAVENEWPARTISIPASI":
                handleSavePartisipasi();break;
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
        for (let i = 0; i < endevent_jam_field.length; i++) {
            numbera = parseInt(endevent_jam_field[i]);
            if (isNaN(numbera)) {
                if (endevent_jam_field[i] != ":") {
                    endevent_jam_field = "";
                }
            }
        }
        
        for (let i = 0; i < mindeposit_field.length; i++) {
            numbera = parseInt(mindeposit_field[i]);
            if (isNaN(numbera)) {
                mindeposit_field = 0;
            }
        }
        for (let i = 0; i < qty_partisipasi_field.length; i++) {
            numbera = parseInt(qty_partisipasi_field[i]);
            if (isNaN(numbera)) {
                qty_partisipasi_field = 0;
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
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DATE</th>
                                <th NOWRAP width="3%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">D</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">INFO</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MIN DEPOSIT</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CASH IN</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">INFORMASI</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterMember as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                NewData("Edit",rec.home_id,rec.home_idwebsite, rec.home_name,
                                                rec.home_start, rec.home_end,
                                                rec.home_mindeposit,
                                                rec.home_status,
                                                rec.home_create,rec.home_update);
                                            }} 
                                            class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i 
                                            on:click={() => {
                                                ListPartisipasi(rec.home_id,rec.home_idwebsite,rec.home_websiteagen,rec.home_name,rec.home_mindeposit,rec.home_status);
                                            }} 
                                            class="bi bi-person-badge"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {rec.home_status}
                                        </span>
                                    </td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        START : {rec.home_start}<br />
                                        END : {rec.home_end}
                                    </td>
                                    <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.home_durationcss}">{rec.home_duration}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        WEBSITE AGEN : {rec.home_websiteagen}<br />
                                        {rec.home_name}
                                    </td>
                                    <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">
                                        {new Intl.NumberFormat().format(rec.home_mindeposit)}
                                    </td>
                                    <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">
                                        {new Intl.NumberFormat().format(rec.home_money_in)}
                                    </td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        CREATE : {rec.home_create}<br />
                                        UPDATE : {rec.home_update}
                                    </td>
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
    modal_body_css="height:550px;overflow-y: scroll;"
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
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Min Deposit</label>
            <input
                bind:value={mindeposit_field}
                on:keyup={handleKeyboard_format}
                class="form-control required"
                type="text"
                style="text-align:right;"
                placeholder=""/>
            <div id="emailHelp" class="form-text" style="color:blue;text-align: right;font-size: 11px;">
                {new Intl.NumberFormat().format(mindeposit_field)}
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
        {#if flag_buttonsave}
        <Button
            on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
        
	</slot:template>
</Modal>

<Modal
	modal_id="modallistpartisipasi"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="List Partisipasi - {nmwebsite_global}"
    modal_body_css="height:600px;overflow-y: scroll;"
    modal_footer_css="padding:2px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="alert alert-primary" role="alert">
            {nmwevent_global}
        </div>
        <ul class="nav nav-tabs" role="tablist">
            <li on:click={() => {
                    TabPartisipasi("LISTALL");
                }} class="nav-item" role="presentation">
                <a class="nav-link active" data-bs-toggle="tab" href="#homepanel" aria-selected="true" role="tab">ListAll</a>
            </li>
            <li on:click={() => {
                    TabPartisipasi("MEMBER");
                }} class="nav-item" role="presentation">
                <a class="nav-link" data-bs-toggle="tab" href="#grouppanel" aria-selected="false" tabindex="-1" role="tab">Member</a>
            </li>
            <li on:click={() => {
                    TabPartisipasi("WINNER");
                }} class="nav-item" role="presentation">
                <a class="nav-link" data-bs-toggle="tab" href="#winnerpannel" aria-selected="false" tabindex="-1" role="tab">Winner</a>
            </li>
        </ul>
        <div id="myTabContent" class="tab-content">
            <div class="tab-pane fade show active" id="homepanel" role="tabpanel">
                <div class="col-lg-12" style="padding: 5px;">
                    <input
                        bind:value={searchListpartisipasi}
                        on:keypress={handleKeyboard_checkenter}
                        type="text"
                        class="form-control"
                        placeholder="Search Event"
                        aria-label="Search"/>
                </div>
                <table class="table table-striped ">
                    <thead>
                        <tr>
                            <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                            <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                            <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                            <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">VOUCHER</th>
                            <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPOSIT</th>
                            <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">WINNER</th>
                            <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREATE</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each filterListpartisipasi as rec}
                        <tr>
                            <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_no}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_username}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                <a href="https://wa.me/{rec.eventdetail_phone}" target="_blank">
                                    {rec.eventdetail_phone}
                                </a>
                            </td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_voucher}</td>
                            <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.eventdetail_deposit)}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_status}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_create}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="tab-pane fade" id="grouppanel" role="tabpanel">
                <table class="table table-striped ">
                    <thead>
                        <tr>
                            <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                            <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                            <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                            <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTAL</th>
                            <th NOWRAP width="*" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPOSIT</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each listmembergroup_db as rec}
                        <tr>
                            <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.eventdetailgroup_no}</td>
                            <td  on:click={() => {
                                ListMemberAgenVoucher(rec.eventdetailgroup_idmember,rec.eventdetailgroup_username,"");
                            }} NOWRAP style="cursor:pointer;text-decoration:underline;text-align: left;vertical-align: top;font-size: {table_body_font};"> {rec.eventdetailgroup_username}</td>
                            <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                <a href="https://wa.me/{rec.eventdetailgroup_phone}" target="_blank">
                                    {rec.eventdetailgroup_phone}
                                </a>
                            </td>
                            <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.eventdetailgroup_voucher)}</td>
                            <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.eventdetailgroup_deposit)}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            <div class="tab-pane fade" id="winnerpannel" role="tabpanel">
                <div class="row">
                    <div class="col-sm-6">
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 1 - {prize1_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize1_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize1_winner_flag}'
                                    type="text"
                                    placeholder="Prize 1"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_1");
                                    }} 
                                    disabled='{prize1_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 2 - {prize2_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize2_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize2_winner_flag}'
                                    type="text"
                                    placeholder="Prize 2"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_2");
                                    }} 
                                    disabled='{prize2_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 3 - {prize3_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize3_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize3_winner_flag}'
                                    type="text"
                                    placeholder="Prize 3"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_3");
                                    }} 
                                    disabled='{prize3_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 4 - {prize4_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize4_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize4_winner_flag}'
                                    type="text"
                                    placeholder="Prize 4"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_4");
                                    }} 
                                    disabled='{prize4_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 5 - {prize5_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize5_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize5_winner_flag}'
                                    type="text"
                                    placeholder="Prize 5"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_5");
                                    }} 
                                    disabled='{prize5_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 6 - {prize6_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize6_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize6_winner_flag}'
                                    type="text"
                                    placeholder="Prize 6"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_6");
                                    }} 
                                    disabled='{prize6_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 7 - {prize7_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize7_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize7_winner_flag}'
                                    type="text"
                                    placeholder="Prize 7"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_7");
                                    }} 
                                    disabled='{prize7_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 8 - {prize8_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize8_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize8_winner_flag}'
                                    type="text"
                                    placeholder="Prize 8"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_8");
                                    }} 
                                    disabled='{prize8_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 9 - {prize9_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize9_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize9_winner_flag}'
                                    type="text"
                                    placeholder="Prize 9"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_9");
                                    }} 
                                    disabled='{prize9_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 10 - {prize10_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize10_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize10_winner_flag}'
                                    type="text"
                                    placeholder="Prize 10"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_10");
                                    }} 
                                    disabled='{prize10_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-6">
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 11 - {prize11_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize11_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize11_winner_flag}'
                                    type="text"
                                    placeholder="Prize 11"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_11");
                                    }} 
                                    disabled='{prize11_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 12 - {prize12_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize12_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize12_winner_flag}'
                                    type="text"
                                    placeholder="Prize 12"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_12");
                                    }} 
                                    disabled='{prize12_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 13 - {prize13_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize13_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize13_winner_flag}'
                                    type="text"
                                    placeholder="Prize 13"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_13");
                                    }} 
                                    disabled='{prize13_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 14 - {prize14_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize14_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize14_winner_flag}'
                                    type="text"
                                    placeholder="Prize 14"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_14");
                                    }} 
                                    disabled='{prize14_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 15 - {prize15_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize15_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize15_winner_flag}'
                                    type="text"
                                    placeholder="Prize 15"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_15");
                                    }} 
                                    disabled='{prize15_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 16 - {prize16_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize16_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize16_winner_flag}'
                                    type="text"
                                    placeholder="Prize 16"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_16");
                                    }} 
                                    disabled='{prize16_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 17 - {prize17_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize17_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize17_winner_flag}'
                                    type="text"
                                    placeholder="Prize 17"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_17");
                                    }} 
                                    disabled='{prize17_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 18 - {prize18_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize18_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize18_winner_flag}'
                                    type="text"
                                    placeholder="Prize 18"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_18");
                                    }} 
                                    disabled='{prize18_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 19 - {prize19_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize19_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize19_winner_flag}'
                                    type="text"
                                    placeholder="Prize 19"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_19");
                                    }} 
                                    disabled='{prize19_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Prize 20 - {prize20_username}</label>
                            <div class="input-group mb-3">
                                <Input
                                    bind:value={prize20_winner}
                                    on:keyup={handleKeyboard_format}    
                                    disabled='{prize20_winner_flag}'
                                    type="text"
                                    placeholder="Prize 20"/>
                                <button
                                    on:click={() => {
                                        handleMemberPopup("PRIZE_20");
                                    }} 
                                    disabled='{prize20_winner_save_flag}' 
                                    type="button" class="btn btn-info">Info</button>
                            </div>
                        </div>
                    </div>
                </div>
                
                
            </div>
        </div>
        
	</slot:template>
	<slot:template slot="footer">
        {#if panel_footer_listall}
            <div style="font-size: 12px;">
                SUBTOTAL : <span style="color:blue;font-weight:bold;">{new Intl.NumberFormat().format(listpartisipasi_total)}</span>
            </div>
            {#if flag_buttonsave}
                <Button
                    on:click={callFunction}
                    button_function="FORM_PARTISIPASI"
                    button_title="New Partisipasi"
                    button_css="btn-warning"/>
            {/if}
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalentrypartisipasi"
	modal_size="modal-dialog-centered"
	modal_title="New Partisipasi"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Member</label>
            <div class="input-group">
                <input 
                    bind:value={idmemberagen_partisipasi_field}
                    type="hidden" >
                <input
                    disabled
                    bind:value={username_partisipasi_field} 
                    type="text" class="form-control required" placeholder="Username">
                <button
                    on:click={() => {
                        ListMemberAgen();
                    }}  
                    class="btn btn-info" type="button">
                    <i class="bi bi-search"></i>
                </button>
                
            </div>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Qty</label>
            <input
                bind:value={qty_partisipasi_field}
                on:keyup={handleKeyboard_format} 
				maxlength="3" 
                class="form-control required"
                type="text"
                style="text-align:right;"
                placeholder=""/>
            <div id="emailHelp" class="form-text" style="color:blue;text-align: right;font-size: 11px;">
                {new Intl.NumberFormat().format(qty_partisipasi_field)}
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
            on:click={callFunction}
            button_function="SAVENEWPARTISIPASI"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modallistmemberagen"
	modal_size="modal-dialog-centered"
	modal_title="List Member Agen"
    modal_body_css="height:400px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                </tr>
            </thead>
            <tbody>
                {#each listmemberagen_db as rec}
                <tr 
                    on:click={() => {
                        InsertPartisipasi(rec.memberagen_id,rec.memberagen_username+" - "+rec.memberagen_phone+" - "+rec.memberagen_name);
                    }}  
                    style="text-decoration: underline;color:biru;cursor: pointer;">
                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_no}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_username}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_phone}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.memberagen_name}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">

        <Button
            on:click={callFunction}
            button_function="FORM_PARTISIPASI"
            button_title="New Partisipasi"
            button_css="btn-warning"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modallistmemberagenvoucher"
	modal_size="modal-dialog-centered"
	modal_title="List Voucher : {username_global}"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">VOUCHER</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREATE</th>
                </tr>
            </thead>
            <tbody>
                {#each listpartisipasivoucher_db as rec}
                <tr>
                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_no}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_voucher}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_create}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
	</slot:template>
</Modal>


<Modal
	modal_id="modallistmembervoucher"
	modal_size="modal-dialog-centered"
	modal_title="List Member Agen"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                    <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTAL</th>
                    <th NOWRAP width="*" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPOSIT</th>
                </tr>
            </thead>
            <tbody>
                {#each listmembergroup_db as rec}
                <tr>
                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.eventdetailgroup_no}</td>
                    <td  on:click={() => {
                        ListMemberAgenVoucher(rec.eventdetailgroup_idmember,rec.eventdetailgroup_username,"WINNER");
                    }} NOWRAP style="cursor:pointer;text-decoration:underline;text-align: left;vertical-align: top;font-size: {table_body_font};"> {rec.eventdetailgroup_username}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                        <a href="https://wa.me/{rec.eventdetailgroup_phone}" target="_blank">
                            {rec.eventdetailgroup_phone}
                        </a>
                    </td>
                    <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.eventdetailgroup_voucher)}</td>
                    <td  NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.eventdetailgroup_deposit)}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">

        <Button
            on:click={callFunction}
            button_function="FORM_PARTISIPASI"
            button_title="New Partisipasi"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modallistmemberagenvoucherwinner"
	modal_size="modal-dialog-centered"
	modal_title="List Voucher : {username_global}"
    modal_body_css="height:500px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">VOUCHER</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREATE</th>
                </tr>
            </thead>
            <tbody>
                {#each listpartisipasivoucher_db as rec}
                <tr>
                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_no}</td>
                    <td  on:click={() => {
                        InsertFirebase(rec.eventdetail_voucher,rec.eventdetail_id);
                    }} NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font}; cursor:pointer;text-decoration:underline;">{rec.eventdetail_voucher}</td>
                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.eventdetail_create}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
	</slot:template>
</Modal>

<Modal
	modal_id="modalviewvoucher"
	modal_size="modal-dialog-centered"
	modal_title="VOUCHER"
    modal_body_css="height:400px;overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <ul style="list-style: none;font-size:12px;">
            {#each listvoucher_view as rec}
                <li>{rec.event_voucher}</li>
            {/each}
        </ul>
	</slot:template>
	<slot:template slot="footer">
	</slot:template>
</Modal>