import React, { useState, useEffect } from 'react'
import { useParams, useHistory } from 'react-router-dom'
import { startCase, camelCase } from 'lodash'
import { useToaster, useConfirmationDialog, Text, Color } from '@harness/uicore'
import { Intent } from '@blueprintjs/core'
import { useStrings } from 'framework/strings'
import { useGetExecution, useDeleteExecution, useUpdateExecution } from 'services/pm'
import { Settings } from '../../components/Settings/Settings'
import routes from 'RouteDefinitions'

interface PathProps {
  pipeline: string
  execution: string
}

interface ExecutionProps {
  name?: string
  desc?: string
}

export const ExecutionSettings = () => {
  const history = useHistory()
  const { getString } = useStrings()
  const { showError, showSuccess } = useToaster()
  const { pipeline, execution } = useParams<PathProps>()

  const [name, setName] = useState<string | undefined>('')
  const [desc, setDesc] = useState<string | undefined>('')

  const { data, loading, error, refetch } = useGetExecution({ pipeline, execution })
  const { mutate: deleteExecution } = useDeleteExecution({ pipeline })
  const { mutate: updateExecution } = useUpdateExecution({ pipeline, execution })
  const title = `${startCase(camelCase(data?.name!.replace(/-/g, ' ')))} Settings`

  useEffect(() => {
    if (data) {
      setName(data.name)
      setDesc(data.desc)
    }
  }, [data])

  const handleUpdate = async ({ name, desc }: ExecutionProps) => {
    try {
      await updateExecution({ name, desc })
      showSuccess(getString('common.itemUpdated'))
      refetch()
    } catch (err) {
      showError(`Error: ${err}`)
      console.error(err)
    }
  }

  const handleDelete = async () => {
    try {
      await deleteExecution(execution)
      history.push(routes.toPipeline({ pipeline }))
    } catch (err) {
      showError(`Error: ${err}`)
      console.error(err)
    }
  }

  const { openDialog } = useConfirmationDialog({
    titleText: getString('common.delete'),
    contentText: <Text color={Color.GREY_800}>Are you sure you want to delete this?</Text>,
    confirmButtonText: getString('common.delete'),
    cancelButtonText: getString('common.cancel'),
    intent: Intent.DANGER,
    buttonIntent: Intent.DANGER,
    onCloseDialog: async (isConfirmed: boolean) => {
      if (isConfirmed) {
        try {
          await handleDelete()
          showSuccess(getString('common.itemDeleted'))
          refetch()
        } catch (err) {
          showError(`Error: ${JSON.stringify(err)}`)
          console.error({ err })
        }
      }
    }
  })

  const handleSubmit = (data: ExecutionProps): void => {
    handleUpdate(data)
  }

  return (
    <Settings
      name={name}
      desc={desc}
      handleDelete={openDialog}
      loading={loading}
      handleSubmit={handleSubmit}
      refetch={refetch}
      title={title}
      error={error}
    />
  )
}