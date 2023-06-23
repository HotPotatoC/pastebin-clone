type ButtonProps = {
  children: React.ReactNode
  onClick?: () => void
  type?: 'button' | 'submit' | 'reset'
  variant?: 'primary' | 'secondary'
}

const Button = ({
  children,
  onClick,
  type = 'button',
  variant = 'primary',
}: ButtonProps) => {
  const baseClass =
    'w-full inline-flex justify-center items-center rounded-2xl border border-transparent px-16 py-5 text-xl font-fira-code font-medium focus:outline-none'

  const variantClass = {
    primary: 'bg-primary text-white duration-100 hover:bg-primary-darken',
    secondary: 'bg-white text-primary',
  }

  return (
    <button
      type={type}
      className={`${baseClass} ${variantClass[variant]}`}
      onClick={onClick}
    >
      {children}
    </button>
  )
}

export default Button
