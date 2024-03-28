import { debounce } from 'lodash-es';
import { useCallback, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Search20Regular } from '@fluentui/react-icons';
import { Combobox, Option } from '@fluentui/react-components';

function useSearch() {
  const [options, setOptions] = useState<any[]>([]);
  const navigate = useNavigate();
  const [searchText, setSearchText] = useState('');

  // eslint-disable-next-line react-hooks/exhaustive-deps
  const getResources = useCallback(
    debounce(async (val: string) => {
      console.log(val);
    }, 300),
    // debounce(async (value: string) => {
    //   if (!value) return;
    //   try {
    //     const { data } = await service('/api/anonymous/search', {
    //       method: 'POST',
    //     })({ data: { text: value } });
    //     const resources = data.resources.map((resource: any) => ({
    //       ...resource,
    //       url: resourceTypeToUrl[
    //         resource.resourceType as ResourceType
    //       ]?.replace(':resourceId', resource.resourceId),
    //     }));
    //     setOptions(resources);
    //   } catch (err: any) {
    //     console.log(err);
    //   } finally {
    //     //
    //   }
    // }, 300)
    []
  );

  const onOptionClick = (url: string) => {
    setSearchText('');
    navigate(url);
  };

  const onSearchTextChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchText(e.target.value);
    getResources(e.target.value);
  };

  return {
    options,
    searchText,
    onOptionClick,
    onSearchTextChange,
  };
}

function Search() {
  const { options, searchText, onOptionClick, onSearchTextChange } =
    useSearch();

  return (
    <div className='w-8 sm:w-96'>
      <Combobox
        value={searchText}
        onChange={onSearchTextChange}
        expandIcon={<Search20Regular />}
        style={{ width: '100%' }}
        placeholder='Search anything on the platform'
        onOptionSelect={(_, data) => onOptionClick(data.optionValue || '')}
      >
        {options.map((option) => (
          <Option key={option.id} value={option.url}>
            {option.name}
          </Option>
        ))}
      </Combobox>
    </div>
  );
}

export default Search;
