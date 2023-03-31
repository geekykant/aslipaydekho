<template>
    <n-card :title="'Compensation Post #' + postId">
        <n-form ref="form" v-model="form" :rules="rules">
            <n-grid x-gap="12" :cols="3" justify="start">
                <n-gi>
                    <n-form-item required label="Date of the Offer" prop="dateOfOffer">
                        <n-date-picker v-model="form.dateOfOffer"></n-date-picker>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item required label="Years of Experience (eg. 2.3)" prop="yearsOfExperience">
                        <n-input-number placeholder="2.3" clearable :precision="1" :show-button="false" v-model="form.yearsOfExperience">
                            <template #suffix>years</template>
                        </n-input-number>
                    </n-form-item>
                </n-gi>
            </n-grid>

            <n-divider title-placement="left">Educational Details</n-divider>
            <n-radio-group v-model:value="Unknown" name="radiogroup">
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
            

            <n-grid x-gap="12" :cols="3">
                <n-gi>
                    <n-form-item label="College Name" prop="collegeName">
                        <n-input placeholder="IIT, NIT" v-model:value="form.collegeName" />
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="Course Program" prop="courseProgram">
                        <n-input placeholder="B.Tech, M.Tech" v-model:value="form.courseProgram" />
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="Course Subject" prop="courseSubject">
                        <n-input placeholder="CSE, ECE" v-model:value="form.courseSubject" />
                    </n-form-item>
                </n-gi>
            </n-grid>

            <n-divider title-placement="left">Offer Letter Details</n-divider>
            <n-grid x-gap="12" :cols="3" justify="space-between">
                <n-gi>
                    
                </n-gi>
            </n-grid>
            <n-grid x-gap="12" :cols="2">
                <n-gi>
                    <n-card title="Prior Company Details">
                        <n-form-item label="Previous Company Name" prop="previousCompanyName">
                            <n-input v-model:value="form.previousCompanyName"/>
                        </n-form-item>
                        <n-form-item label="Previous Job Title" prop="previousJobTitle">
                            <n-input v-model:value="form.previousJobTitle"/>
                        </n-form-item>
                        <n-form-item label="Previous TC (Salary + Bonus + Stock)" prop="previousSalary">
                            <n-input-number :show-button="false" round :precision="0" placeholder="9,00,000" v-model="form.salary">
                                <template #suffix>₹</template>
                            </n-input-number>
                        </n-form-item>
                        <n-form-item label="Previous Job Location" prop="location">
                            <n-input v-model="form.previousJoblocation"/>
                        </n-form-item>
                    </n-card>
                </n-gi>
                <n-gi>
                    <n-card title="Offer Company Details">
                        <n-form-item required label="Company" prop="company">
                            <n-input v-model="form.company"/>
                        </n-form-item>
                        <n-form-item required label="Offer Job Title/Level" prop="titleLevel">
                            <n-input v-model="form.titleLevel"/>
                        </n-form-item>
                        <n-form-item required label="New TC (Salary + Bonus + Stock)" prop="totalComp">
                            <n-input-number :show-button="false" clearable :precision="0" placeholder="15,00,000" v-model="form.totalComp">
                                <template #suffix>₹</template>
                            </n-input-number>
                        </n-form-item>
                        <n-form-item label="Offer Location" prop="location">
                            <n-input v-model="form.location"/>
                        </n-form-item>
                    </n-card>
                </n-gi>
            </n-grid>

            <br>
            <br>

            <n-grid x-gap="12" :cols="3" justify="space-between">
                <n-gi>
                    <n-form-item required label="Base Salary" prop="baseSalary">
                        <n-input-number :show-button="false" round placeholder="14,00,000" v-model="form.salary">
                            <template #suffix>₹</template>
                        </n-input-number>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="Relocation/Signing Bonus" prop="relocationBonus">
                        <n-input-number :show-button="false" v-model:value="form.relocationBonus">
                            <template #suffix>₹</template>
                        </n-input-number>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="Stock bonus" prop="stockBonus">
                        <n-input-number :show-button="false" v-model:value="form.stockBonus">
                            <template #suffix>$</template>
                        </n-input-number>
                    </n-form-item>
                </n-gi>
            </n-grid>

            <n-form-item label="Performance bonus (in %)" prop="performanceBonus">
                <n-input-number :show-button="false" round :precision="0" v-model:value="form.performanceBonus">
                    <template #suffix>%</template>
                </n-input-number>
            </n-form-item>
            <n-form-item label="Extra Remarks" prop="extraRemarks">
                <n-input placeholder="Extra Remarks" v-model:value="form.extraRemarks" type="textarea" maxlength="400" show-count />
            </n-form-item>
            
        </n-form>
        <br>
        <n-button type="primary" @click="submitForm">Submit</n-button>
    </n-card>
</template>

<script>
export default {
    props: {
        data: String | null,
        postUrl: String | null,
    },
    computed: {
        postId() {
            if(this.postUrl == null) return "";
            return this.postUrl.split("/").pop();
        },
    },
    data() {
        return {
            form: {
            },
            rules: {
                education: [{ required: true, message: "Please input education", trigger: "blur" }],
                yearsOfExperience: [{ required: true, message: "Please input years of experience", trigger: "blur" }],
                priorExperience: [{ required: true, message: "Please input prior experience", trigger: "blur" }],
                dateOfOffer: [{ required: true, message: "Please select date of offer", trigger: "blur" }],
                company: [{ required: true, message: "Please input company", trigger: "blur" }],
                titleLevel: [{ required: true, message: "Please input title/level", trigger: "blur" }],
                location: [{ required: true, message: "Please input location", trigger: "blur" }],
                salary: [{ required: true, message: "Please input salary", trigger: "blur" }],
                relocationBonus: [{ required: true, message: "Please input relocation/signing bonus", trigger: "blur" }],
                stockBonus: [{ required: true, message: "Please input stock bonus", trigger: "blur" }],
                bonus: [{ required: true, message: "Please input bonus", trigger: "blur" }],
                totalComp: [{ required: true, message: "Please input total comp", trigger: "blur" }],
                extraRemarks: [{ required: true, message: "Please input extra remarks", trigger: "blur" }]
            },
            value: ref(null),
            songs: [
                {
                    value: "Tier 1",
                    label: "Tier 1"
                },
                {
                    value: "Tier 2",
                    label: "Tier 2"
                },
                {
                    value: "Tier 3",
                    label: "Tier 3"
                },
                {
                    value: "Unknown",
                    label: "Unknown"
                }
            ].map((s) => {
                s.value = s.value.toLowerCase();
                return s;
            })
        };
    },
}
</script>


<style>
    .form-div{
        border: 1px solid;
        padding: 20px 10px;
    }
</style>