<template>
    <div class="container" >

        <!-- <Wallet v-if="page === 'A'" />
        <Recent v-else-if="page === 'B'"/>
        <Settings v-else="page === 'C'"/> -->
        <p v-if="page === 'A'"> A {{ page }} </p>
        <p v-if="page === 'B'"> B {{ page }} </p>
        <p v-if="page === 'C'"> C {{ page }} </p>
        <div class="row fixed-bottom fq iconUp" style="background:black">
            <div class="col" @click="page = 'A'">
                <font-awesome-icon  style="color:white"  icon="fas fa-solid fa-wallet fa-lg" /> 
            </div>
            <div class="col" @click="page = 'B'">
                <font-awesome-icon style="color:white"  icon="fas fa-solid fa-timeline fa-lg" /> 
            </div>
            <div class="col" @click="page = 'C'">
                <font-awesome-icon style="color:white"  icon="fas fa-solid fa-gears fa-lg" /> 
            </div>
        </div>
    </div>
        <!-- <h2 class="amount"> $54.00 </h2> -->
</template>
<script>
import Recent from './Recent'
import Settings from './Settings'
import Wallet from './Wallet'

export default {
    components: {
        Recent,
        Settings,
        Wallet
    },
    data() {
        return {
            active: true,
            page: 'A',
            list: "example.com",
            icons: {
                active: "images/icon-48x48.png",
                inactive: "images/icon-48x48-off.png"
            },
            
        };
    },
    created() {
        chrome.storage.sync.get(["toggleSitesActive", "toggleSitesList"], (result) => {
            this.active = result.toggleSitesActive;
            this.list = result.toggleSitesList;
        });
    },
    methods: {
        setActive(active) {
            this.active = active;
            chrome.storage.sync.set({
                toggleSitesActive: active
            }, () => { });
            chrome.browserAction.setIcon({
                path: this.icons[active ? "active" : "inactive"]
            });
        },
        saveList() {
            chrome.storage.sync.set({
                toggleSitesList: this.list
            }, () => { });
        },
                setPage (num) {
            this.page = num
        }

    }
}
</script>