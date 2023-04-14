<template>
    <n-card style="margin: 0px">
        <n-form ref="formRef" :model="formValue" :rules="formRules">
            <n-space justify="space-between">
                <n-h3 prefix="bar" type="info">Compensation Post #{{ postId }}</n-h3>
                <nuxt-link :to="postUrl" target="_blank"><n-icon size="24"><LinkOutline/></n-icon></nuxt-link>
            </n-space>
            <div>
                <label>Choose Offer Type: </label>
                <n-radio-group v-model:value="formValue.offerType" prop="offerType">
                    <n-radio-button
                        v-for="offerType in offerTypes"
                        :key="offerType.value"
                        :value="offerType.value"
                        :label="offerType.label"
                    />
                </n-radio-group>
            </div>
            <br>
            <n-grid x-gap="12" :cols="3" justify="start">
                <n-gi>
                    <n-form-item required label="Date of the Offer" prop="dateOfOffer">
                        <n-date-picker format="dd-MM-yyyy" v-model:value="formValue.dateOfOffer"></n-date-picker>
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item required label="Years of Experience (eg. 2.3)" prop="yearsOfExperience">
                        <n-input-number placeholder="2.3" clearable :precision="1" :show-button="false" v-model:value="formValue.yearsOfExperience">
                            <template #suffix>years</template>
                        </n-input-number>
                    </n-form-item>
                </n-gi>
            </n-grid>

            <n-tabs type="line" animated>
                <n-tab-pane name="educationDetails" tab="Educational Details">
                    <n-radio-group v-model:value="formValue.collegeTier" name="radiogroup" size="medium">
                        <n-space>
                            <label>College Tier:</label>
                            <n-radio-button
                                v-for="tier in collegeTiers"
                                :key="tier.value"
                                :value="tier.value"
                                :label="tier.label"
                            />
                        </n-space>
                    </n-radio-group>
                    <br>
                    <br>
                    <n-grid x-gap="12" :cols="3">
                        <n-gi>
                            <n-form-item label="College Name" prop="collegeName">
                                <n-input placeholder="IIT, NIT, IIIT" v-model:value="formValue.collegeName" />
                            </n-form-item>
                        </n-gi>
                        <n-gi>
                            <n-form-item label="Course Program" prop="courseProgram">
                                <n-input placeholder="B.Tech, M.Tech" v-model:value="formValue.courseProgram" />
                            </n-form-item>
                        </n-gi>
                        <n-gi>
                            <n-form-item label="Course Subject" prop="courseSubject">
                                <n-input placeholder="CSE, ECE" v-model:value="formValue.courseSubject" />
                            </n-form-item>
                        </n-gi>
                    </n-grid>
                </n-tab-pane>
                <n-tab-pane name="newOfferDetails" tab="Company Details">
                    <n-grid x-gap="12" :cols="formValue.yearsOfExperience == 0 ? 1 : 2">
                        <n-gi :key="formValue.yearsOfExperience" v-show="formValue.yearsOfExperience != 0">
                            <n-card title="Prior Company Details">
                                <n-form-item label="Previous Company" prop="previousCompanyName">
                                    <n-input v-model:value="formValue.previousCompany.companyName"/>
                                </n-form-item>
                                <n-form-item label="Previous Job Title" prop="previousJobTitle">
                                    <n-input v-model:value="formValue.previousCompany.jobTitle"/>
                                </n-form-item>
                                <n-form-item label="Previous Job Location" prop="location">
                                    <n-input v-model:value="formValue.previousCompany.jobLocation"/>
                                </n-form-item>
                                <n-form-item label="Previous TC (Salary + Bonus + Stock)" prop="previousSalary">
                                    <n-input-number :show-button="false" round :precision="0" placeholder="9,00,000" v-model:value="formValue.previousCompany.totalComp">
                                        <template #suffix>₹</template>
                                    </n-input-number>
                                </n-form-item>
                            </n-card>
                        </n-gi>
                        <n-gi>
                            <n-card title="Offer Company Details">
                                <n-form-item required label="Company" prop="company">
                                    <n-input v-model:value="formValue.newCompany.companyName"/>
                                </n-form-item>
                                <n-form-item required label="Offer Job Title/Level" prop="titleLevel">
                                    <n-input v-model:value="formValue.newCompany.jobTitle"/>
                                </n-form-item>
                                <n-form-item label="Offer Location" prop="location">
                                    <n-input v-model:value="formValue.newCompany.jobLocation"/>
                                </n-form-item>
                                <n-form-item required label="New TC (Salary + Bonus + Stock)" prop="totalComp">
                                    <n-input-number :show-button="false" clearable :precision="0" placeholder="15,00,000" v-model:value="formValue.newCompany.totalComp">
                                        <template #suffix>₹</template>
                                    </n-input-number>
                                </n-form-item>
                            </n-card>
                        </n-gi>
                    </n-grid>
                </n-tab-pane>
                <n-tab-pane name="moreDetails" tab="More Pay Details">
                    <n-grid x-gap="12" :cols="3" justify="space-between">
                    <n-gi>
                        <n-form-item required label="Base Salary" prop="baseSalary">
                            <n-input-number :show-button="false" round placeholder="14,00,000" v-model:value="formValue.newCompany.baseSalary">
                                <template #suffix>₹</template>
                            </n-input-number>
                        </n-form-item>
                    </n-gi>
                    <n-gi>
                        <n-form-item label="Relocation/Signing Bonus" prop="relocationBonus">
                            <n-input-number :show-button="false" v-model:value="formValue.newCompany.relocationBonus">
                                <template #suffix>₹</template>
                            </n-input-number>
                        </n-form-item>
                    </n-gi>
                    <n-gi>
                        <n-form-item label="Stock bonus" prop="stockBonus">
                            <n-input-number :show-button="false" v-model:value="formValue.newCompany.stockBonus">
                                <template #suffix>$</template>
                            </n-input-number>
                        </n-form-item>
                    </n-gi>
                </n-grid>

                <n-form-item label="Performance bonus" prop="performanceBonus">
                    <n-input-number :show-button="false" round :precision="0" v-model:value="formValue.newCompany.performanceBonus">
                        <template #suffix>% per year</template>
                    </n-input-number>
                </n-form-item>
                <n-form-item label="Extra Remarks" prop="extraRemarks">
                    <n-input placeholder="Extra Remarks" v-model:value="formValue.newCompany.extraRemarks" type="textarea" maxlength="400" show-count />
                </n-form-item>
                </n-tab-pane>
            </n-tabs>
        </n-form>

        <!-- <pre>{{ JSON.stringify(formValue, null, 2) }}</pre> -->
        
        <br>
        <n-button type="primary" @click="submitForm">Submit</n-button>
    </n-card>
</template>

<script setup lang="ts">
import { LinkOutline } from '@vicons/ionicons5';
import { FormInst, useMessage } from 'naive-ui';

const { postFormData } = defineProps({
    postFormData: JSON | null
});

const formRef = ref<FormInst | null>(null);
const formValue = ref({
    offerType: "fulltime",
    collegeTier: "TierThree",
    dateOfOffer: new Date(postFormData.post.creationDate),
    previousCompany: {},
    newCompany: {}
});

const postUrl = postFormData.url;
const postId = postFormData.post.id;

const collegeTiers = [
    {
        label: "Tier 1/2",
        value: "TierOneOrTwo"
    },
    {
        label: "Tier 3",
        value: "TierThree"
    },
];

const offerTypes = [
    {
        value: "internship",
        label: "Internship"
    },
    {
        value: "fulltime",
        label: "Fulltime"
    },
];

const message = useMessage()

const submitForm = (e: MouseEvent) => {
    e.preventDefault()
    const messageReactive = message.loading('Verifying', {
        duration: 1
    });

    if(formValue.value.offerType == null)
        message.error("Offer Type is not valid");
    else if(formValue.value.dateOfOffer == null)
        message.error("Date of Offer is not valid");
    else if(formValue.value.yearsOfExperience == null)
        message.error("Years of Experience is not valid");
    else if(formValue.value.collegeTier == null)
        message.error("College Tier is not valid");
    else if(formValue.value.newCompany?.companyName == null)
        message.error("New Company Name is not valid");
    else if(formValue.value.newCompany?.jobTitle == null)
        message.error("New Job Title is not valid");
    else if(formValue.value.newCompany?.totalComp == null)
        message.error("New Job Total Compensation is not valid");
    else if(formValue.value.newCompany?.baseSalary == null)
        message.error("New Company Base Salary is not valid");
    else
        message.success("All good!")
    
    messageReactive.destroy();
}
</script>


<style>
    .form-div{
        border: 1px solid;
        padding: 20px 10px;
    }
</style>