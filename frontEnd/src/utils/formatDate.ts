export const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    const options: Intl.DateTimeFormatOptions  = { 
      day: '2-digit', 
      month: '2-digit', 
      year: 'numeric', 
      hour: '2-digit', 
      minute: '2-digit', 
      second: '2-digit', 
      hour12: false, 
      timeZone: 'Europe/Paris' 
    };
    
    const formattedDate = new Intl.DateTimeFormat('fr-FR', options).format(date);
    return formattedDate;
  }