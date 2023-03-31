<template>
    <div class="content-wrapper">
        <n-row :gutter="12">
            <n-col :span="15" offset="1">
                <n-skeleton v-if="pending" :width="146" :sharp="false" size="medium" />
                <PostForm :data="postFormData" :postUrl="postUrl" v-else />
            </n-col>
            <n-col :span="7">
                <n-skeleton v-if="pending" text :repeat="5" sharp="false"/>
                <n-skeleton v-if="pending" text style="width: 60%" sharp="false"/>
                <UnParsedPostContent :data="unParsedPostData" v-else/>
            </n-col>
        </n-row>
    </div>
</template>


<script type="ts" setup>
const unParsedPostData = ref();
const postFormData = ref(); 
const postUrl = ref();

const { data, pending} =  await useLazyFetch('/api/fetch-post', {
    async onResponse({request, response, err}) {
        if(response.ok){
            console.log(response._data)
            data.value = response._data;
            unParsedPostData.value = data.value.post.content,
            postFormData.value = data.value.offerLetter,
            postUrl.value = data.value.url
        }
        else{
            //TODO: handle error case
        }
    }
})
</script>

<style>
</style>