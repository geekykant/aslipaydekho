<template>
    <n-alert title="Original Post" type="info">
        <div v-html="formattedData"/>
    </n-alert>
</template>

<script>
export default {
    props: {
        unParsedPostData: String | null
    },
    computed: {
        formattedData() {
            if(this.unParsedPostData == null) return "";
            return this.unParsedPostData
                .replace(/\\n/g, "<br>") // Convert '\n' to '<br>' for HTML line breaks
                .replace(/^\*/gm, "- "); // Replace markdown bullet points with dash
        }
    },
}
</script>