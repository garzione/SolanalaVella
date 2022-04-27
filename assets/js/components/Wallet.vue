<template>
    <div>
        <div class="row" >
            <h1 class="title">' la Vella</h1>
        </div>
        <div id="hoverText" class="row">
            <div class="row fq">
                <div v-on:click="copyAndSwapColor" > {{ truncateAddress() }} </div>
            </div>

            <div class="row fq">
                <div id="toolT" class="col1" > {{ !copyIt ? 'Copy to clipboard' : 'Copied!' }} </div>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            copyIt: false,
            wallet: "LnK5ScoNvxpEJJGA5tjNEkmv5tUhSECbtpm6pQBkA13"
        }

    },
    created() {
    },
    methods: {
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

}
}
</script>