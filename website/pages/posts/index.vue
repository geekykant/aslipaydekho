<template>
    <div class="content-wrapper">
        <n-data-table size="small" :columns="columns" :data="data" :pagination="pagination" striped />
    </div>
</template>

<script lang="ts">
import { NButton, DataTableColumns, NDivider } from 'naive-ui'

type RowData = {
    key: number
    name: string
    age: number
    address: string
    tags: string[]
}

const createColumns = ({
    sendMail
}: {
    sendMail: (rowData: RowData) => void
}): DataTableColumns<RowData> => {
    return [
        {
            type: 'expand',
            expandable: (rowData) => rowData.name !== 'Jim Green',
            renderExpand: (rowData) => {
                return `${rowData.name} is a good guy.`
            }
        },
        {
            title: 'Post Id',
            key: 'key',
            width: '10%'
        },
        {
            title: 'Company',
            key: 'name',
            width: '10%'
        },
        {
            title: 'Title',
            key: 'name'
        },
        {
            title: 'YOE',
            key: 'age'
        },
        {
            title: 'Action',
            key: 'actions',
            render(row) {
                return [
                    h(
                        NButton,
                        {
                            type: 'info',
                            size: 'small',
                            onClick: () => sendMail(row)
                        },
                        { default: () => 'Edit' }
                    ),
                    h(NDivider,
                    {
                        vertical: true
                    }
                    ),
                    h(
                        NButton,
                        {
                            type: 'error',
                            size: 'small',
                            onClick: () => sendMail(row)
                        },
                        { default: () => 'Delete' }
                    )
                ]
            },
            width: '15%'
        }
    ]
}

var createData = (): RowData[] => {
    let data : RowData[] = []

    for (let i = 101; i < 150; i++) {
        data.push({
            key: i,
            name: 'John Brown' + i,
            age: 32,
            address: 'New York No. 1 Lake Park',
            tags: ['amazon', 'microsoft']
        });
    }

    return data;
}

export default defineComponent({
    setup() {
        // const message = useMessage()
        return {
            data: createData(),
            columns: createColumns({
                sendMail(rowData) {
                    // message.info('send mail to ' + rowData.name)
                }
            }),
            pagination: {
                pageSize: 10
            }
        }
    }
})
</script>