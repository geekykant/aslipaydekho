<template>
    <n-card title="Compensation Post #3298223">
        <n-form ref="form" v-model="form" :rules="rules">
            <n-divider title-placement="left">Educational Details</n-divider>
            <n-form-item label="College Name (eg. IIT, NIT)" prop="collegeName">
                <n-input placeholder="IIT" v-model="form.collegeName"></n-input>
            </n-form-item>
            <n-form-item label="Course Program (eg. B.Tech, M.Tech)" prop="courseProgram">
                <n-input placeholder="B.Tech" v-model="form.courseProgram"></n-input>
            </n-form-item>
            <n-form-item label="Course Subject (eg. CSE, ECE)" prop="courseSubject">
                <n-input placeholder="CSE" v-model="form.courseSubject"></n-input>
            </n-form-item>
            <n-radio-group v-model:value="value" name="radiogroup">
                <n-space>
                <n-radio
                    v-for="song in songs"
                    :key="song.value"
                    :value="song.value"
                    :label="song.label"
                />
                </n-space>
            </n-radio-group>
            <br>
            <br>
            <n-divider title-placement="left">Basic Details</n-divider>
            <n-form-item required label="Years of Experience (eg. 2.3 - for 2 year and 3 months)" prop="yearsOfExperience">
                <n-input-number placeholder="2.3" clearable :precision="1" :show-button="false" v-model="form.yearsOfExperience">
                    <template #suffix></template>
                </n-input-number>
            </n-form-item>
            
            <n-divider title-placement="left">Offer Letter Details</n-divider>
            <n-grid x-gap="12" :cols="2">
                <n-gi>
                    <n-card title="Prior Company Details">
                        <n-form-item label="Previous Company Name" prop="previousCompanyName">
                            <n-input v-model="form.previousCompanyName"></n-input>
                        </n-form-item>
                        <n-form-item label="Previous Job Title" prop="previousJobTitle">
                            <n-input v-model="form.previousJobTitle"></n-input>
                        </n-form-item>
                        <n-form-item label="Previous TC (Salary + Bonus + Stock)" prop="previousSalary">
                            <n-input-number :show-button="false" clearable :precision="0" placeholder="9,00,000" v-model="form.salary">
                                <template #suffix>₹</template>
                            </n-input-number>
                        </n-form-item>
                    </n-card>
                </n-gi>
                <n-gi>
                    <n-card title="Offer Company Details">
                        <n-form-item required label="Company" prop="company">
                            <n-input v-model="form.company"></n-input>
                        </n-form-item>
                        <n-form-item required label="Offer Job Title/Level" prop="titleLevel">
                            <n-input v-model="form.titleLevel"></n-input>
                        </n-form-item>
                        <n-form-item required label="New TC (Salary + Bonus + Stock)" prop="totalComp">
                            <n-input-number :show-button="false" clearable :precision="0" placeholder="15,00,000" v-model="form.totalComp">
                                <template #suffix>₹</template>
                            </n-input-number>
                        </n-form-item>
                    </n-card>
                </n-gi>
            </n-grid>

            <br>
            <br>

            <n-form-item required label="Date of the Offer" prop="dateOfOffer">
                <n-date-picker v-model="form.dateOfOffer"></n-date-picker>
            </n-form-item>
            <n-form-item label="Offer Location" prop="location">
                <n-input v-model="form.location"></n-input>
            </n-form-item>
            <n-form-item required label="Base Salary" prop="baseSalary">
                <n-input-number :show-button="false" round placeholder="14,00,000" v-model="form.salary">
                    <template #suffix>₹</template>
                </n-input-number>
            </n-form-item>
            <n-form-item label="Relocation/Signing Bonus" prop="relocationBonus">
                <n-input-number :show-button="false" v-model="form.relocationBonus">
                    <template #suffix>₹</template>
                </n-input-number>
            </n-form-item>
            <n-form-item label="Stock bonus" prop="stockBonus">
                <n-input-number :show-button="false" v-model="form.stockBonus">
                    <template #suffix>$</template>
                </n-input-number>
            </n-form-item>
            <n-form-item label="Performance Bonus" prop="bonus">
                <n-input :show-button="false" v-model="form.performanceBonus">
                </n-input>
            </n-form-item>
            <n-form-item label="Extra Remarks" prop="extraRemarks">
                <n-input placeholder="Extra Remarks" v-model="form.extraRemarks" type="textarea" maxlength="30" show-count />
            </n-form-item>
            
        </n-form>
        <br>
        <n-button type="primary" @click="submitForm">Submit</n-button>
    </n-card>
</template>

<script>
export default {
    data() {
        return {
            form: {
                education: '',
                yearsOfExperience: '',
                priorExperience: '',
                dateOfOffer: '',
                company: '',
                titleLevel: '',
                location: '',
                salary: '',
                relocationBonus: '',
                stockBonus: '',
                bonus: '',
                totalComp: '',
                extraRemarks: ''
            },
            rules: {
                education: [{ required: true, message: 'Please input education', trigger: 'blur' }],
                yearsOfExperience: [{ required: true, message: 'Please input years of experience', trigger: 'blur' }],
                priorExperience: [{ required: true, message: 'Please input prior experience', trigger: 'blur' }],
                dateOfOffer: [{ required: true, message: 'Please select date of offer', trigger: 'blur' }],
                company: [{ required: true, message: 'Please input company', trigger: 'blur' }],
                titleLevel: [{ required: true, message: 'Please input title/level', trigger: 'blur' }],
                location: [{ required: true, message: 'Please input location', trigger: 'blur' }],
                salary: [{ required: true, message: 'Please input salary', trigger: 'blur' }],
                relocationBonus: [{ required: true, message: 'Please input relocation/signing bonus', trigger: 'blur' }],
                stockBonus: [{ required: true, message: 'Please input stock bonus', trigger: 'blur' }],
                bonus: [{ required: true, message: 'Please input bonus', trigger: 'blur' }],
                totalComp: [{ required: true, message: 'Please input total comp', trigger: 'blur' }],
                extraRemarks: [{ required: true, message: 'Please input extra remarks', trigger: 'blur' }]
            },
            value: ref(null),
            songs: [
                {
                    value: "Tier 1",
                    label: "Tier 1"
                },
                {
                    value: 'Tier 2',
                    label: 'Tier 2'
                },
                {
                    value: 'Tier 3',
                    label: 'Tier 3'
                },
                {
                    value: 'Unknown',
                    label: 'Unknown'
                }
            ].map((s) => {
                s.value = s.value.toLowerCase()
                return s
            }),
        }
    }
}
</script>


<style>
    .form-div{
        border: 1px solid;
        padding: 20px 10px;
    }
</style>