<template>
    <n-dropdown :options="options">
        <!-- <n-button>User profile</n-button> -->
        <n-avatar
          round
          :size="24"
          src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQMIIpLfirLTkOj6XzRCTq1_2-K-_3ePlIipycBUuspyg"
        />
    </n-dropdown>
</template>

<script lang="ts">
import { NIcon } from 'naive-ui'
import {
    PersonCircleOutline as UserIcon,
    Pencil as EditIcon,
    LogOutOutline as LogoutIcon
} from '@vicons/ionicons5'

const renderIcon = (icon: Component) => {
    return () => {
        return h(NIcon, null, {
            default: () => h(icon)
        })
    }
}

export default defineComponent({
    setup() {
        return {
            options: [
                {
                    label: 'Edit Profile',
                    key: 'editProfile',
                    icon: renderIcon(EditIcon)
                },
                {
                    label: 'Logout',
                    key: 'logout',
                    icon: renderIcon(LogoutIcon)
                }
            ]
        }
    }
})
</script>