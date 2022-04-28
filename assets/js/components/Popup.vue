<template>
    <div class="container" >
        <!-- WALLLET -->
        <div v-if="page === 'A'">
            <div class="row" >
                <h1 class="title">' la Vella</h1>
            </div>
            <div id="hoverText" class="row fq">
                <div class="row fq">
                    <div v-on:click="copyAndSwapColor" > {{ truncateAddress() }} </div>
                </div>

                <div class="row fq">
                    <div id="toolT" class="col1" > {{ !copyIt ? 'Copy to clipboard' : 'Copied!' }} </div>
                </div>
            </div>
            <div class="row fq">
                <p class="fontUp"> {{ "$" + accountBalanceUSD }} </p>
            </div>

            <div class="row sendButton fixed-bottom">
                <button @click="page = 'D'" type="button" class="btn btn-dark btn-lg ">Send</button>
            </div>
        </div>

        <!-- RECENT -->
        <div v-else-if="page === 'B'">
            <div class="row">
                <h1 class="title"> Recent Activity </h1>
            </div>

            <div class="row recentBuffer">
                <div class="alert alert-dark" role="alert">
                    Received $54.60
                </div>
            </div>
            <div class="row recentBuffer">
                <div class="alert alert-dark " role="alert">
                    Sent $67.87
                </div>
            </div>
        </div>

        <!-- SETTINGS -->
        <div v-else-if="page === 'C'">
            <div class="row">
                <h1 class="title"> Settings </h1>
            </div>

            <div class="row recentBuffer">
                <div class="alert alert-dark" role="alert">
                    Wallet Alias
                </div>
            </div>
            <div class="row recentBuffer">
                <div class="alert alert-dark " role="alert">
                    Address Book
                </div>
            </div>
            <div class="row recentBuffer">
                <div class="alert alert-dark " role="alert">
                    Change Password
                </div>
            </div>
        </div>

        <!-- SEND -->
        <div v-else-if="page === 'D'">
            <div class="row">
                <h1 class="title"> Send </h1>
            </div>
            <form @submit="onSubmit">
                <div class="row">
                    <div class="input-group mb-3">
                        <span id="units" @click="currency = !currency" class="input-group-text curs">{{ currency ? "$" : SOLsymbol }}</span>
                        <input v-model="amount" min="1" type="number" placeholder="amount" class="form-control" aria-label="Dollar amount (with dot and two decimal places)">
                    </div>
                    <div class="input-group mb-3">
                        <input v-model="address" placeholder="SOL address" type="text" class="form-control" aria-label="address">
                    </div>
                    <div class="input-group mb-3 fq">
                        <button type="button" class="btn btn-light"> Send </button>
                    </div>
                </div>
            </form>
        </div>

        <div class="row fixed-bottom fq iconUp" style="background:black">
            <div class="col curs" @click="page = 'A'">
                <font-awesome-icon  style="color:white"  icon="fas fa-solid fa-wallet fa-lg" /> 
            </div>
            <div class="col curs" @click="page = 'B'">
                <font-awesome-icon style="color:white"  icon="fas fa-solid fa-timeline fa-lg" /> 
            </div>
            <div class="col curs" @click="page = 'C'">
                <font-awesome-icon style="color:white"  icon="fas fa-solid fa-gears fa-lg" /> 
            </div>
        </div>
    </div>
       
</template>
<script>

export default {
    data() {
        return {
            active: true,
            page: 'A',
            list: "example.com",
            icons: {
                active: "images/icon-48x48.png",
                inactive: "images/icon-48x48-off.png"
            },
            copyIt: false,
            wallet: "LnK5ScoNvxpEJJGA5tjNEkmv5tUhSECbtpm6pQBkA13",
            accountBalanceUSD: 42.32,
            accountBalanceSOL: 2.42,
            SOLconst: 9678,
            currency: 0,
            address: "",
            amount: 0,
            SOLsymbol: String.fromCharCode(9678)
            
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
        },
        copyAndSwapColor() {
            navigator.clipboard.writeText(this.wallet);
            this.copyIt = true;
            const el = document.getElementById("wallet");
            el.classList.remove("col1");
            el.classList.add("col2");
            setTimeout(() => {
                this.copyIt = false;
                el.classList.remove("col2");
                el.classList.add("col1");
            }, 3000);
        },
        truncateAddress() {
            return "(" + this.wallet.slice(0, 3) + "..." + this.wallet.slice(this.wallet.length - 3, this.wallet.length) + ")";
        },
        async onSubmit(e) {
            e.preventDefault()
            const units = Document.getElementById("units")
            if (units === this.SOLsymbol) {
                if (this.accountBalanceSOL < this.amount) {
                    alert("Insufficient funds")
                    return
                } else {
                    //submit to blockchain
                }
            }
            else if (units === "$") {
                if(this.accountBalanceUSD < this.amount) {
                    alert("Insufficient funds")
                    return
                } else {
                    //convert and submit to blockchain
                }
            } else {
                alert("Something went wrong")
                return
            } 

        }

    }
}
</script>