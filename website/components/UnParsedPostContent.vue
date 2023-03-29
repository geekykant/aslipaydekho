<template>
    <n-card title="Post Title">
        <div v-html="formattedPayload"/>
    </n-card>
</template>

<script>
export default {
    data() {
        return {
            payload: `Education: MS in Computer Science\\nYears of Experience: 2 years\\nPrior Experience: NA\\nCompany: Samsung\\nTitle/Level: Associate Engineer\\nLocation: Seattle\\nSalary: 96000\\nRelocation: 15000\\nSigning Bonus: 0\\nStock bonus: 0\\nBonus: Performance-based bonus up to 20% of salary every year\\nTotal comp (Salary + Bonus + Stock): ~$120K\\nBenefits: 15 paid vacation days + 10 paid U.S. holidays  + 401k/+ covered health/dental/vision\\nOther details: No Negotiated.`,
        }
    },

    computed: {
        formattedPayload() {
            return this.payload.replace(/\\n/g, "<br>") // Convert '\n' to '<br>' for HTML line breaks
                .replace(/^\*/gm, "- ") // Replace markdown bullet points with dash
        }
    }
}
</script>