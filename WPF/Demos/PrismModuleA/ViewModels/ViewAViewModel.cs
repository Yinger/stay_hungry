using Prism.Commands;
using Prism.Events;
using Prism.Mvvm;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace PrismModuleA.ViewModels
{
    public class ViewAViewModel : BindableBase
    {
        private string _message;
        private string title;
        private readonly IEventAggregator eventAggregator;

        public string Title
        {
            get { return title; }
            set
            {
                title = value;
                RaisePropertyChanged();
            }
        }

        public string Message
        {
            get { return _message; }
            set { SetProperty(ref _message, value); }
        }

        public DelegateCommand OpenCommand { get; private set; }
        public DelegateCommand OpenCommand2 { get; private set; }

        public CompositeCommand OpenAll { get; private set; }

        public DelegateCommand EventCommand { get; private set; }
        public DelegateCommand SendCommand { get; private set; }

        public ViewAViewModel(IEventAggregator eventAggregator)
        {
            Message = "View A from your Prism Module";
            Title = "View AAA";
            OpenCommand = new DelegateCommand(() =>
            {
                Title += "Prism\r\n";
            });
            OpenCommand2 = new DelegateCommand(() =>
            {
                Title += "Demo\r\n";
            });

            OpenAll = new CompositeCommand();
            OpenAll.RegisterCommand(OpenCommand);
            OpenAll.RegisterCommand(OpenCommand2);

            this.eventAggregator = eventAggregator;
            EventCommand = new DelegateCommand(() =>
            {
                eventAggregator.GetEvent<MessageEvent>().Subscribe(OnMessageReceived, ThreadOption.PublisherThread, false, msg =>
                {
                    if (msg.Equals("Hello"))
                        return true;
                    else
                        return false;
                });
            });
            SendCommand = new DelegateCommand(() =>
            {
                eventAggregator.GetEvent<MessageEvent>().Publish("Aloha");
                //eventAggregator.GetEvent<MessageEvent>().Unsubscribe(OnMessageReceived);
            });
        }

        public void OnMessageReceived(string message)
        {
            Title = Title + "\r\n" + message;
        }
    }

    public class MessageEvent : PubSubEvent<string>
    { }
}