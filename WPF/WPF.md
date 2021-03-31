# WPF 布局容器

- StackPanel: 水平或垂直排列元素、Orientation 属性分别: Horizontal / Vertical
- WrapPanel : 水平或垂直排列元素、针对剩余空间不足会进行换行或换列进行排列
- DockPanel : 根据容器的边界、元素进行 Dock.Top、Left、Right、Bottom 设置
- Grid : 类似 table 表格、可灵活设置行列并放置控件元素、比较常用
- UniformGrid : 指定行和列的数量, 均分有限的容器空间
- Canvas : 使用固定的坐标设置元素的位置、不具备锚定停靠等功能。

# 控件

![](/WPF/wpf.png)

# Prism

- https://github.com/PrismLibrary/Prism
- https://github.com/PrismLibrary/Prism-Samples-Wpf

## Region

目的: 弱化模块与模块之间的耦合关系

### 1. 定义区域

#### 定义方式一(XML)

```xml
<Grid>
    <ContentControl prism:RegionManager.RegionName="ContentRegion" />
</Grid>
```

```C#
    public class MainWindowViewModel : BindableBase
    {
        private readonly IRegionManager regionManager;

        public MainWindowViewModel(IRegionManager regionManager)
        {
            this.regionManager = regionManager;
            regionManager.RegisterViewWithRegion("ContentRegion", typeof(ViewA));
        }
    }
```

#### 定义方式二(Code)

```xml
<Grid>
    <ContentControl x:Name="ctr" />
</Grid>
```

```C#
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();

            RegionManager.SetRegionName(ctr, "ContentRegion");
        }
    }
```

### 2. 维护区域集合

### 3.提供对区域的访问

### 4.合成视图

### 5.区域导航

### RegionAdapter

Prism 内置了几个区域适配器，所以我们可以在 ContentControl 当中定义区域，实际可以在任何元素上定义区域，如果定义的范围不在官方提供的默认适配器当中，则会引发异常。

- ContentControlRegionAdapter
- ItemsControlRegionAdapter
- SelectorRegionAdapter
  - ComboBox
  - ListBox
  - Ribbon
  - TabControl

```xml
<Grid>
    <StackPanel prism:RegionManager.RegionName="ContentRegion" />
</Grid>
```

```C#
using Prism.Regions;
using System.Windows;
using System.Windows.Controls;

namespace PrismDemo
{
    public class StackPanelRegionAdapter : RegionAdapterBase<StackPanel>
    {
        public StackPanelRegionAdapter(IRegionBehaviorFactory regionBehaviorFactory) : base(regionBehaviorFactory)
        {
        }

        protected override void Adapt(IRegion region, StackPanel regionTarget)
        {
            region.Views.CollectionChanged += (s, e) =>
            {
                if (e.Action == System.Collections.Specialized.NotifyCollectionChangedAction.Add)
                {
                    foreach (FrameworkElement item in e.NewItems)
                    {
                        regionTarget.Children.Add(item);
                    }
                }
            };
        }

        protected override IRegion CreateRegion()
        {
            return new Region();
        }
    }
}
```

```C#
    public partial class App
    {
        protected override Window CreateShell()
        {
            return Container.Resolve<MainWindow>();
        }

        protected override void RegisterTypes(IContainerRegistry containerRegistry)
        {
        }

        protected override void ConfigureRegionAdapterMappings(RegionAdapterMappings regionAdapterMappings)
        {
            base.ConfigureRegionAdapterMappings(regionAdapterMappings);

            regionAdapterMappings.RegisterMapping(typeof(StackPanel), Container.Resolve<StackPanelRegionAdapter>());
        }
    }
```

## Module

# MVVM

- ICommand
- INotifyPropertyChanged

| ↓ 機能/→Framework | Prism            | Mvvmlight     | Micorsoft.Toolkit.Mvvm |
| ----------------- | ---------------- | ------------- | ---------------------- |
| Binding           | BindableBase     | ViewModelBase | ObservableObject       |
| Command           | DelegateCommand  | RelayCommand  | Async/RelayCommand     |
| Aggregator        | IEventAggregator | IMessager     | IMessager              |
| Module 化         | ✔                | ×             | ×                      |
| Container         | ✔                | ×             | ×                      |
| Ioc               | ✔                | ×             | ×                      |
| navigation        | ✔                | ×             | ×                      |
| Dialog(service)   | ✔                | ×             | ×                      |
