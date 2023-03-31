<template>
    <n-alert title="Original Post" type="info">
        <div v-html="formattedData"/>
    </n-alert>
</template>

<script>
export default {
    props: {
        data: String | null
    },
    computed: {
        formattedData() {
            if(this.data == null) return "";
            return this.data
                .replace(/\\n/g, "<br>") // Convert '\n' to '<br>' for HTML line breaks
                .replace(/^\*/gm, "- "); // Replace markdown bullet points with dash
        }
    },
}
</script>