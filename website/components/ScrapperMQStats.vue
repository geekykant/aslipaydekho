<template>
    <n-card id="published-stats" title="Cron Broker Statistics">
        <n-row>
            <n-col :span="3">
                <n-statistic label="Published Posts" :value="0">
                    <template #prefix>
                        <n-icon><Rocket/></n-icon>
                    </template>
                    <template #suffix>/ 100</template>
                </n-statistic>
            </n-col>
            <n-col :span="3">
                <n-statistic label="Active Users">1,234,123</n-statistic>
            </n-col>
        </n-row>
    </n-card>
</template>

<script lang="ts">
import { Rocket } from '@vicons/ionicons5'

export default defineComponent({
    components: {
        Rocket
    }
})
</script>

<style>
#published-stats{
    background-color: white;
}
</style>