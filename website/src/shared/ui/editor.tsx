import { useRef } from 'react'
import useAutosizeTextArea from '../hooks/useAutosizeTextArea'

type EditorProps = {
  viewOnly?: boolean
  value?: string
  onChange?: (value: string) => void
}

const Editor = ({ viewOnly, value = '', onChange }: EditorProps) => {
  const textAreaRef = useRef<HTMLTextAreaElement>(null)
  useAutosizeTextArea(textAreaRef.current, value)

  return (
    <div className="my-10 flex h-full w-full items-start space-x-6 rounded-2xl bg-grey px-6 py-5 font-fira-code text-base font-medium focus:outline-none md:px-10 md:text-lg">
      <div className="flex h-full flex-col items-end">
        {value?.split('\n').map((_, idx) => (
          <div key={idx} className="font-fira-code font-medium opacity-50">
            {idx + 1}
          </div>
        ))}
      </div>
      <div className="flex flex-grow">
        <textarea
          className="h-full w-full flex-grow resize-none bg-transparent focus:outline-none"
          onChange={(e) => onChange?.(e.target.value)}
          placeholder="Start typing here..."
          ref={textAreaRef}
          rows={12}
          value={value}
          readOnly={viewOnly}
        ></textarea>
      </div>
    </div>
  )
}

export default Editor
