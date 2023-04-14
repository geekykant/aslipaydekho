<template>
    <div class="content-wrapper">
        <n-message-provider>
        <n-row :gutter="12">
            <n-col :span="15" offset="1">
                <n-skeleton v-if="pending" :width="146" :sharp="false" size="medium" />
                <PostForm :postFormData="postFormData" v-else />
            </n-col>
            <n-col :span="7">
                <n-skeleton v-if="pending" text :repeat="5" sharp="false"/>
                <n-skeleton v-if="pending" text style="width: 60%" sharp="false"/>
                <UnParsedPostContent :unParsedPostData="unParsedPostData" v-else/>
            </n-col>
        </n-row>
        </n-message-provider>
    </div>
</template>

<script type="ts" setup>
const unParsedPostData = ref();
const postFormData = ref(); 

const { pending } =  await useLazyFetch('/api/fetch-post', {
    async onResponse({request, response, err}) {
        if(response.ok){
            var resData = response._data;
            postFormData.value = resData;
            unParsedPostData.value = resData.post.content;
        }
        else{
            //TODO: handle error case
        }
    }
})
</script>

<style>
</style>