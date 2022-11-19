const timeStringToBrDateTimeString = (strDate: string) => {
  const dt = new Date(strDate)
  return `${dt.toLocaleDateString('pt-BR')} ${dt.toLocaleTimeString('pt-BR')}`
}

const dateToBrDateTimeString = (date: Date) => {
  return `${date.toLocaleDateString('pt-BR')} ${date.toLocaleTimeString('pt-BR')}`
}

const timeStringToBrDateString = (strDate: string) => {
  const dt = new Date(strDate)
  return `${dt.toLocaleDateString('pt-BR')}`
}

export { timeStringToBrDateTimeString, dateToBrDateTimeString, timeStringToBrDateString }
